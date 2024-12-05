package core

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/labstack/gommon/log"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/fingo/pkg/types"
	"github.com/sinameshkini/microkit/models"
	"time"
)

// Transfer is an internal API for other modules to do any transaction
func (c *Core) Transfer(ctx context.Context, req endpoint.TransferRequest) (resp *endpoint.TransferResponse, err error) {
	var (
		l *redislock.Lock
	)

	if req.RawAmount+req.FeeAmount != req.TotalAmount {
		return nil, entities.ErrInvalidRequest
	}

	if !req.SkipLock {
		key := fmt.Sprintf("fingo:lock:%s", req.DebitAccountID)

		l, err = c.lock.Obtain(ctx, key, 10*time.Second, nil)
		if err != nil {
			return
		}
	}

	tx := c.repo.NewTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("%+v", r))
			tx.Rollback()
		}
	}()

	txn := &entities.Transaction{
		UserID:      req.UserID,
		OrderID:     req.OrderID,
		Type:        req.Type,
		Amount:      types.Amount(req.TotalAmount),
		Description: req.Description,
	}

	if err = tx.Create(&txn).Error; err != nil {
		tx.Rollback()
		return
	}

	// fee transaction
	if req.FeeAmount != 0 {
		if err = c.repo.Transfer(tx,
			types.Amount(req.FeeAmount),
			txn.ID,
			models.ParseIDf(req.DebitAccountID),
			models.ParseIDf(req.FeeAccountID),
			req.FeeDescription,
		); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// do transfer
	if err = c.repo.Transfer(tx,
		types.Amount(req.RawAmount),
		txn.ID,
		models.ParseIDf(req.DebitAccountID),
		models.ParseIDf(req.CreditAccountID),
		req.Description,
	); err != nil {
		tx.Rollback()
		return
	}

	// commit transaction
	if err = c.repo.CommitTransaction(tx); err != nil {
		return
	}

	if !req.SkipLock && l != nil {
		if err = l.Release(ctx); err != nil {
			return
		}
	}

	return c.GetTransaction(ctx, req.UserID, txn.ID)
}

func (c *Core) GetTransaction(ctx context.Context, userID string, txnID models.SID) (resp *endpoint.TransferResponse, err error) {
	txn, err := c.repo.GetTransaction(ctx, txnID)
	if err != nil {
		return
	}

	resp = txn.ToResponse(userID)
	if resp == nil {
		return nil, entities.ErrNotFound
	}

	return
}

func (c *Core) Reverse(ctx context.Context, req endpoint.ReverseRequest) (resp *endpoint.TransferResponse, err error) {
	var (
		transaction *entities.Transaction
	)

	// get txn by ID
	if transaction, err = c.repo.GetTransaction(ctx, models.ParseIDf(req.TransactionID)); err != nil {
		return
	}

	if req.UserID != transaction.UserID {
		return nil, entities.ErrPermissionDenied
	}

	// do reverse
	tx := c.repo.NewTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("%+v", r))
			tx.Rollback()
		}
	}()

	reverseTxn, err := c.repo.Initial(tx, req.UserID, transaction.OrderID, enums.Reverse, transaction.Amount, enums.Reverse.Label())
	if err != nil {
		tx.Rollback()
		return
	}

	if err = c.repo.Reverse(tx, transaction, reverseTxn.ID); err != nil {
		tx.Rollback()
		return
	}

	// commit transaction
	if err = c.repo.CommitTransaction(tx); err != nil {
		return
	}

	return c.GetTransaction(ctx, reverseTxn.UserID, reverseTxn.ID)
}
