package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
)

func (r *repo) GetPolicies(ctx context.Context, userID, accountID, accountType string) (policies []*entities.Policy, err error) {
	query := r.db.WithContext(ctx).
		Where("is_enable = true")

	if userID != "" {
		query = query.Or("entity_type = 'user' AND entity_id = ?", userID)
	}

	if accountID != "" {
		query = query.Or("entity_type = 'account' AND entity_id = ?", accountID)
	}

	if accountType != "" {
		query = query.Or("entity_type = 'account_type' AND entity_id = ?", accountType)
	}

	if err = query.Order("priority ASC").
		Find(&policies).Error; err != nil {
		return
	}

	return
}
