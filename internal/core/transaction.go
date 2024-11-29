package core

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/labstack/gommon/log"
	"github.com/sinameshkini/fingo/internal/models"
	"time"
)

// Transfer is an internal API for other modules to do any transaction
func (c *Core) Transfer(ctx context.Context, req models.TransferRequest) (resp *models.TransferResponse, err error) {
	var (
		l *redislock.Lock
	)

	if req.RawAmount+req.FeeAmount != req.TotalAmount {
		return nil, models.ErrInvalidRequest
	}

	if !req.SkipLock {
		key := c.lock.SetKeyPrefix(req.DebitAccountID)

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

	txn := &models.Transaction{
		OrderID:     req.OrderID,
		Type:        req.Type,
		Amount:      req.TotalAmount,
		Description: req.Description,
	}

	if err = tx.Create(&txn).Error; err != nil {
		tx.Rollback()
		return
	}

	// fee transaction
	if req.FeeAmount != 0 {
		if err = c.repo.Transfer(tx,
			req.FeeAmount,
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
		req.RawAmount,
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

func (c *Core) GetTransaction(ctx context.Context, userID string, txnID models.ID) (resp *models.TransferResponse, err error) {
	txn, err := c.repo.GetTransaction(ctx, txnID)
	if err != nil {
		return
	}

	resp = txn.ToResponse(userID)
	if resp == nil {
		return nil, models.ErrNotFound
	}

	return
}

//func (s *svc) Reverse(ctx context.Context, orderID, transactionID string) (resp *models.TxnResponse, err error) {
//	var (
//		transaction *entities.Transaction
//	)
//
//	// get txn by ID
//	if transaction, err = c.repo.GetByID(ctx, transactionID); err != nil {
//		return
//	}
//
//	// check reverse exist
//	if transaction.Reverse != nil {
//		return nil, enums.ErrAlreadyExist
//	}
//
//	// do reverse
//	tx := c.repo.NewTransaction(ctx)
//	defer func() {
//		if r := recover(); r != nil {
//			logstash.Get().Error(context.Background()).Commit(fmt.Sprintf("%+v", r))
//			tx.Rollback()
//		}
//	}()
//
//	reverseTxn, err := c.repo.Initial(ctx, tx, orderID, enums.Reverse, transaction.Amount, enums.Reverse.Label())
//	if err != nil {
//		tx.Rollback()
//		return
//	}
//
//	if err = c.repo.Reverse(ctx, tx, transaction, reverseTxn.ID); err != nil {
//		tx.Rollback()
//		return
//	}
//
//	// commit transaction
//	if err = c.repo.CommitTransaction(ctx, tx); err != nil {
//		return
//	}
//
//	resp = &models.TxnResponse{
//		TransactionID: reverseTxn.ID,
//		OrderID:       reverseTxn.OrderID,
//		Description:   reverseTxn.Description,
//		Message:       "یازگشت تراکنش با موفقیت انجام شد",
//	}
//
//	return
//}
