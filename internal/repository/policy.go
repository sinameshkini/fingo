package repository

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/microkit/models"
)

func (r *repo) FetchPolicies(ctx context.Context, req endpoint.FetchPoliciesRequest) (resp []*entities.Policy, meta *models.PaginationResponse, err error) {
	query := r.db.WithContext(ctx).Model(&entities.Policy{})

	total, err := models.GetCount(query)
	if err != nil {
		return
	}

	query = req.PaginationRequest.ToQuery(query)

	if err = query.Find(&resp).Error; err != nil {
		return
	}

	meta = models.MakePaginationResponse(total, req.Page, req.PerPage)

	return
}

func (r *repo) CreatePolicy(ctx context.Context, req entities.Policy) (resp *entities.Policy, err error) {
	if err = r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *repo) UpdatePolicy(ctx context.Context, policyID models.SID, req entities.Policy) (resp *entities.Policy, err error) {
	if err = r.db.WithContext(ctx).Model(&entities.Policy{}).Where("id = ?", policyID).Updates(req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *repo) DeletePolicy(ctx context.Context, policyID models.SID) (err error) {
	if err = r.db.WithContext(ctx).Where("id = ?", policyID).Delete(&entities.Policy{}).Error; err != nil {
		return
	}
	return
}

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
