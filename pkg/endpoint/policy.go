package endpoint

import "github.com/sinameshkini/microkit/models"

type FetchPoliciesRequest struct {
	models.Request
}

type GetSettingsRequest struct {
	UserID        string `query:"user_id"`
	AccountID     string `query:"account_id"`
	AccountTypeID string `query:"account_type_id"`
}
