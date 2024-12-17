package core

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/endpoint"
	"github.com/sinameshkini/fingo/pkg/enums"
	"github.com/sinameshkini/microkit/models"
)

func (c *Core) FetchPolicies(ctx context.Context, req endpoint.FetchPoliciesRequest) (resp []*entities.Policy, meta *models.PaginationResponse, err error) {
	return c.repo.FetchPolicies(ctx, req)
}

func (c *Core) CreatePolicy(ctx context.Context, req entities.Policy) (resp *entities.Policy, err error) {
	return c.repo.CreatePolicy(ctx, req)
}

func (c *Core) UpdatePolicy(ctx context.Context, policyID models.SID, req entities.Policy) (resp *entities.Policy, err error) {
	return c.repo.UpdatePolicy(ctx, policyID, req)
}

func (c *Core) DeletePolicy(ctx context.Context, policyID models.SID) (err error) {
	return c.repo.DeletePolicy(ctx, policyID)
}

func (c *Core) GetSettings(ctx context.Context, req endpoint.GetSettingsRequest) (settings *entities.Settings, err error) {
	var (
		policies []*entities.Policy
		sp       = entities.SettingsP{
			Limits: make(map[string]entities.LimitsP),
			//Limits: &entities.LimitsP{
			//	NumberOfAccounts: make(map[string]uint),
			//},
			Codes: make(map[enums.ProcessCode]entities.CodeP),
		}
	)

	if policies, err = c.repo.GetPolicies(ctx, req.UserID, req.AccountID, req.AccountTypeID); err != nil {
		//s.logger.Error(ctx).
		//	With(log2.FieldMethodName, "GetSettings").
		//	With(log2.FieldUserID, req.UserID).
		//	With("account_id", req.AccountID).
		//	With(log2.FieldMethodInput, fmt.Sprintf("account_type: %s, user_groups: %v", req.AccountTypeID, req.UserGroups)).
		//	Commit("permission denied")
		return
	}

	for _, p := range policies {
		sp = merge(sp, p.Settings)
	}

	if settings, err = validate(sp); err != nil {
		//s.logger.Error(ctx).
		//	With(log2.FieldMethodName, "GetSettings").
		//	With(log2.FieldUserID, req.UserID).
		//	With("account_id", req.AccountID).
		//	With(log2.FieldMethodInput, fmt.Sprintf("account_type: %s, user_groups: %v", req.AccountTypeID, req.UserGroups)).
		//	With("policies", fmt.Sprintf("%+v", policies)).
		//	Commit("invalid settings")
		return
	}

	return
}

func validate(sp entities.SettingsP) (s *entities.Settings, err error) {
	err = enums.ErrPermissionDenied
	if sp.Limits == nil {
		return
	}

	s = &entities.Settings{
		Limits: make(map[string]entities.Limits),
		Codes:  make(map[enums.ProcessCode]entities.Code),
	}

	for at, l := range sp.Limits {
		if v := l.MinBalance; v == nil {
			return
		}

		if v := l.MaxBalance; v == nil {
			return
		}

		if v := l.NumberOfAccounts; v == nil {
			return
		}

		s.Limits[at] = entities.Limits{
			MinBalance:       *l.MinBalance,
			MaxBalance:       *l.MaxBalance,
			NumberOfAccounts: *l.NumberOfAccounts,
		}
	}

	for pc, c := range sp.Codes {
		if v := c.FeeType; v == nil {
			return
		}

		if v := c.FeeValue; v == nil {
			return
		}

		if v := c.MinAmountPerTransaction; v == nil {
			return
		}

		if v := c.MaxAmountPerTransaction; v == nil {
			return
		}

		if v := c.MaxAmountPerDay; v == nil {
			return
		}

		if v := c.MaxCountPerDay; v == nil {
			return
		}

		s.Codes[pc] = entities.Code{
			FeeType:                 *c.FeeType,
			FeeValue:                *c.FeeValue,
			MinAmountPerTransaction: *c.MinAmountPerTransaction,
			MaxAmountPerTransaction: *c.MaxAmountPerTransaction,
			MaxAmountPerDay:         *c.MaxAmountPerDay,
			MaxCountPerDay:          *c.MaxCountPerDay,
		}
	}

	if sp.DefaultAccountTypeID == nil {
		return
	}

	s.DefaultAccountTypeID = *sp.DefaultAccountTypeID

	err = nil

	return
}

func merge(sp, in entities.SettingsP) entities.SettingsP {
	for at, l := range in.Limits {
		existed, ok := sp.Limits[at]
		if !ok {
			sp.Limits[at] = l
			continue
		}

		if v := l.MinBalance; v != nil {
			existed.MinBalance = v
		}

		if v := l.MaxBalance; v != nil {
			existed.MaxBalance = v
		}

		if v := l.NumberOfAccounts; v != nil {
			existed.NumberOfAccounts = v
		}

		sp.Limits[at] = existed
	}

	for pc, c := range in.Codes {
		existed, ok := sp.Codes[pc]
		if !ok {
			sp.Codes[pc] = c
			continue
		}

		if v := c.FeeType; v != nil {
			existed.FeeType = v
		}

		if v := c.FeeValue; v != nil {
			existed.FeeValue = v
		}

		if v := c.MinAmountPerTransaction; v != nil {
			existed.MinAmountPerTransaction = v
		}

		if v := c.MaxAmountPerTransaction; v != nil {
			existed.MaxAmountPerTransaction = v
		}

		if v := c.MaxAmountPerDay; v != nil {
			existed.MaxAmountPerDay = v
		}

		if v := c.MaxCountPerDay; v != nil {
			existed.MaxCountPerDay = v
		}

		sp.Codes[pc] = existed
	}

	if in.DefaultAccountTypeID != nil {
		sp.DefaultAccountTypeID = in.DefaultAccountTypeID
	}

	return sp
}
