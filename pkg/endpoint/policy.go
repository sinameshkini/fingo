package endpoint

type GetSettingsRequest struct {
	UserID        string `query:"user_id"`
	AccountID     string `query:"account_id"`
	AccountTypeID string `query:"account_type_id"`
}
