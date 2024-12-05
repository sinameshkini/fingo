package core

import (
	"context"
	"github.com/sinameshkini/fingo/internal/repository/entities"
	"github.com/sinameshkini/fingo/pkg/enums"
)

func (c *Core) GetSettings(ctx context.Context, req entities.GetSettingsRequest) (settings *entities.Settings, err error) {
	var (
		policies []*entities.Policy
		sp       = entities.SettingsP{
			Limits: &entities.LimitsP{
				NumberOfAccounts: make(map[string]uint),
			},
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

	if sp.Limits.MaxBalance == nil {
		return
	}

	if sp.Limits.MinBalance == nil {
		return
	}

	s = &entities.Settings{
		Limits: entities.Limits{
			MinBalance:       *sp.Limits.MinBalance,
			MaxBalance:       *sp.Limits.MaxBalance,
			NumberOfAccounts: make(map[string]uint),
		},
		Codes: make(map[enums.ProcessCode]entities.Code),
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

	s.Limits.NumberOfAccounts = sp.Limits.NumberOfAccounts

	err = nil

	return
}

func merge(sp, in entities.SettingsP) entities.SettingsP {
	if in.Limits != nil {
		if in.Limits.MinBalance != nil {
			sp.Limits.MinBalance = in.Limits.MinBalance
		}

		if in.Limits.MaxBalance != nil {
			sp.Limits.MaxBalance = in.Limits.MaxBalance
		}
	}

	for t, n := range in.Limits.NumberOfAccounts {
		sp.Limits.NumberOfAccounts[t] = n
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
