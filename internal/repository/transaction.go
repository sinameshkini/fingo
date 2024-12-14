package repository

import (
	"context"
	"fmt"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
	"gorm.io/gorm"
)

func (r *repo) GetBalance(ctx context.Context, accountID models.SID) (resp models.Amount, err error) {
	if err = r.db.WithContext(ctx).
		Model(&entities.Document{}).
		Where("account_id = ?", accountID).
		Select("COALESCE(sum(amount), 0) as balance").
		Scan(&resp).Error; err != nil {
		return
	}

	return
}

func (r *repo) NewTransaction(
	ctx context.Context,
	userID, orderID, description string,
	txnType enums.TransactionType,
	amount models.Amount,
) (
	txn *entities.Transaction,
	tx *gorm.DB,
	err error,
) {

	tx = r.db.WithContext(ctx).Begin()

	txn = &entities.Transaction{
		UserID:      userID,
		OrderID:     orderID,
		Type:        txnType,
		Amount:      amount,
		Description: description,
	}

	if err = tx.Create(&txn).Error; err != nil {
		return
	}

	return
}

func (r *repo) CommitTransaction(tx *gorm.DB) (err error) {
	return tx.Commit().Error
}

func (r *repo) Transfer(tx *gorm.DB, amount models.Amount, txnID, debID, credID models.SID, comment string) (err error) {
	var (
		debBalance  models.Amount
		credBalance models.Amount
	)
	if err = tx.Model(&entities.Document{}).
		Where("account_id = ?", debID).
		Select("COALESCE(sum(amount), 0)").Scan(&debBalance).Error; err != nil {
		return
	}

	debitDocument := &entities.Document{
		TransactionID:  txnID,
		AccountID:      debID,
		AccountPartyID: credID,
		Type:           enums.Debit,
		Amount:         amount * -1,
		Comment:        comment,
		Balance:        debBalance - amount,
	}

	if err = tx.Create(&debitDocument).Error; err != nil {
		return
	}

	if err = tx.Model(&entities.Document{}).
		Where("account_id = ?", credID).
		Select("COALESCE(sum(amount), 0)").Scan(&credBalance).Error; err != nil {
		return
	}

	creditDocument := &entities.Document{
		TransactionID:  txnID,
		AccountID:      credID,
		AccountPartyID: debID,
		Type:           enums.Credit,
		Amount:         amount,
		Comment:        comment,
		Balance:        credBalance + amount,
	}

	if err = tx.Create(&creditDocument).Error; err != nil {
		return
	}

	return
}

func (r *repo) Reverse(tx *gorm.DB, transaction *entities.Transaction, reverseTxnID models.SID) (err error) {
	for i := len(transaction.Documents) - 1; i >= 0; i-- {
		var (
			balance models.Amount
			d       = transaction.Documents[i]
		)

		if err = tx.Model(&entities.Document{}).
			Where("account_id = ?", d.AccountID).
			Select("COALESCE(sum(amount), 0)").Scan(&balance).
			Error; err != nil {
			return err
		}

		reverseDoc := &entities.Document{
			TransactionID:  reverseTxnID,
			AccountID:      d.AccountID,
			AccountPartyID: d.AccountPartyID,
			Type:           d.Type.Reverse(),
			Amount:         d.Amount * -1,
			Comment:        fmt.Sprintf("اصلاحیه %s", d.Comment),
			Balance:        balance - d.Amount,
		}

		if err = tx.Create(&reverseDoc).Error; err != nil {
			return err
		}
	}

	return
}

func (r *repo) GetTransaction(ctx context.Context, txnID models.SID) (resp *entities.Transaction, err error) {
	if err = r.db.WithContext(ctx).
		Preload("Documents.Account").
		First(&resp, txnID).Error; err != nil {
		return
	}

	return
}

func (r *repo) Inquiry(ctx context.Context, req endpoint.InquiryRequest) (resp []*entities.Transaction, err error) {
	query := r.db.WithContext(ctx).
		Preload("Documents.Account")

	if txnID := models.ParseSIDf(req.TransactionID); txnID != 0 {
		query = query.Where("id = ?", txnID)
	}

	if req.OrderID != "" {
		query = query.Where("order_id = ?", req.OrderID)

	}
	if err = query.Where("user_id = ?", req.UserID).
		Find(&resp).Error; err != nil {
		return
	}

	return
}

func (r *repo) GetHistory(ctx context.Context, req endpoint.HistoryRequest) (resp []*entities.Document, meta *models.PaginationResponse, err error) {
	query := r.db.WithContext(ctx).
		Model(&entities.Document{}).
		Where("account_id = ?", req.AccountID)

	total, err := models.GetCount(query)
	if err != nil {
		return nil, nil, err
	}

	query = req.ToQuery(query)

	if err = query.
		//Joins("join documents on transactions.id = documents.transaction_id").
		//Where("documents.account_id = ?", accountID).
		//Preload("Documents.Account").
		Preload("Account").
		Preload("Transaction").
		Order("id desc").
		Find(&resp).Error; err != nil {
		return
	}

	meta = models.MakePaginationResponse(total, req.Page, req.PerPage)

	return
}

func (r *repo) GetByOrderID(ctx context.Context, userID, orderID string) (resp []*entities.Transaction, err error) {
	if err = r.db.WithContext(ctx).
		Preload("Documents.Account").
		Where("user_ID = ?", userID).
		Where("order_id = ?", orderID).
		First(&resp).Error; err != nil {
		return
	}

	return
}
