package core

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/labstack/gommon/log"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"github.com/sinameshkini/microkit/pkg/utils"
	"time"
)

// Transfer is an internal API for other modules to do any transaction
func (c *Core) Transfer(ctx context.Context, req endpoint.TransactionRequest) (resp *endpoint.TransactionResponse, err error) {
	var (
		locks []*redislock.Lock
	)

	// validate request
	if err = c.validate.Struct(req); err != nil {
		return
	}

	// validate request amounts
	if err = req.ValidateAmount(); err != nil {
		return
	}

	_, err = c.repo.GetByOrderID(ctx, req.UserID, req.OrderID)
	if err == nil {
		return nil, models.ErrAlreadyExist
	}

	txn, tx, err := c.repo.NewTransaction(ctx, req.UserID, req.OrderID, req.Description, req.Type, req.TotalAmount)
	if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("%+v", r))
			tx.Rollback()
		}
	}()

	if c.env.Lock {
		defer func() {
			for _, l := range locks {
				if err = l.Release(ctx); err != nil {
					return
				}
			}
		}()
	}

	for _, transfer := range req.Transfers {
		if c.env.Lock && !transfer.SkipLock {
			key := fmt.Sprintf("fingo:lock:%s", transfer.DebitAccountID)

			l, err := c.lock.Obtain(ctx, key, 10*time.Second, nil)
			if err != nil {
				return nil, enums.ErrToManyRequests
			}

			locks = append(locks, l)
		}

		if err = c.repo.Transfer(
			tx,
			transfer.Amount,
			txn.ID,
			models.ParseSIDf(transfer.DebitAccountID),
			models.ParseSIDf(transfer.CreditAccountID),
			transfer.Comment,
		); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// commit transaction
	if err = c.repo.CommitTransaction(tx); err != nil {
		return
	}

	return c.GetTransaction(ctx, req.UserID, txn.ID)
}

func (c *Core) GetTransaction(ctx context.Context, userID string, txnID models.SID) (resp *endpoint.TransactionResponse, err error) {
	txn, err := c.repo.GetTransaction(ctx, txnID)
	if err != nil {
		return
	}

	resp = txn.ToResponse(userID)
	if resp == nil {
		return nil, enums.ErrNotFound
	}

	return
}

func (c *Core) Reverse(ctx context.Context, req endpoint.ReverseRequest) (resp *endpoint.TransactionResponse, err error) {
	var (
		transaction *entities.Transaction
	)

	// get txn by ID
	if transaction, err = c.repo.GetTransaction(ctx, models.ParseSIDf(req.TransactionID)); err != nil {
		return
	}

	if req.UserID != transaction.UserID {
		return nil, enums.ErrPermissionDenied
	}

	// do reverse
	reverseTxn, tx, err := c.repo.NewTransaction(ctx, req.UserID, transaction.OrderID, req.Description, enums.Reverse, transaction.Amount)
	if err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Error(fmt.Sprintf("%+v", r))
			tx.Rollback()
		}
	}()

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

func (c *Core) History(ctx context.Context, req endpoint.HistoryRequest) (resp *endpoint.HistoryResponse, err error) {
	resp = &endpoint.HistoryResponse{}
	docs, meta, err := c.repo.GetHistory(ctx, req)
	if err != nil {
		return
	}

	for _, doc := range docs {
		resp.Transactions = append(resp.Transactions, doc.ToResponse(req.UserID))
	}

	resp.Meta = meta

	return
}

func (c *Core) Inquiry(ctx context.Context, req endpoint.InquiryRequest) (resp []*endpoint.TransactionResponse, err error) {
	transactions, err := c.repo.Inquiry(ctx, req)
	if err != nil {
		return
	}

	for _, txn := range transactions {
		resp = append(resp, txn.ToResponse(req.UserID))
	}

	if resp == nil {
		return nil, enums.ErrNotFound
	}

	return
}

// Enqueue adds a new transaction to the queue
func (c *Core) Enqueue(req endpoint.TransactionRequest) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.queue <- req
	}()
}

// startWorker starts a single worker to process transactions serially
func (c *Core) startWorker() {
	go func() {
		for req := range c.queue {
			// TODO handle error
			resp, err := c.Transfer(context.Background(), req)
			if err != nil {
				log.Error(fmt.Sprintf("transaction_error: %+v", err))
			}
			utils.PrintJson(resp)
		}
	}()
}

// Stop waits for all transactions to complete and closes the queue
func (c *Core) Stop() {
	c.wg.Wait()
	close(c.queue)
}
