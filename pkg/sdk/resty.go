package sdk

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sinameshkini/fingo/internal/models"
	"github.com/sinameshkini/fingo/pkg/utils"
)

type Client struct {
	rc *resty.Client
}

func New(baseURL string) *Client {
	return &Client{
		rc: resty.New().
			SetDebug(true).
			SetBaseURL(baseURL),
	}
}

func (c *Client) CreateAccount(req models.CreateAccount) (resp *models.AccountResponse, err error) {
	apiResp := &models.Response{}

	r, err := c.rc.R().
		SetBody(&req).
		SetResult(apiResp).
		SetError(apiResp).
		Post("/accounts")
	if err != nil {
		return
	}

	if r.IsError() {
		err = errors.New(r.String())
	}

	if err = utils.JsonAssertion(apiResp.Data, &resp); err != nil {
		return
	}

	return
}

//
//func (c *Client) GetWallet() (resp *GetWalletResponse, err error) {
//	resp = &GetWalletResponse{}
//
//	r, err := c.rc.R().
//		SetResult(resp).
//		Get("")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		return nil, errors.InitID(r.Status())
//	}
//
//	return
//}
//
//
//func (c *Client) CreateWalletVerify(otp string) (resp *CreateWalletVerifyOtpResponse, apiErr *APIError, err error) {
//	resp = &CreateWalletVerifyOtpResponse{}
//	apiErr = &APIError{}
//
//	r, err := c.rc.R().
//		SetBody(map[string]interface{}{"otpCode": otp}).
//		SetResult(resp).
//		SetError(apiErr).
//		Post("/verify-otp")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		err = fmt.Errorf("code: %d, msg: %s", apiErr.Code, apiErr.Message)
//	}
//
//	return
//}
//
//func (c *Client) Deposit(amount string) (claims jwt.MapClaims, apiErr *APIError, err error) {
//	orderResp := &DepositWalletResponse{}
//	apiErr = &APIError{}
//
//	r, err := c.rc.R().
//		SetBody(map[string]interface{}{"amount": amount}).
//		SetResult(orderResp).
//		SetError(apiErr).
//		Post("/deposit")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		err = fmt.Errorf("code: %d, msg: %s", apiErr.Code, apiErr.Message)
//	}
//
//	if orderResp.OrderId == "" {
//		return
//	}
//
//	token, _, err := new(jwt.Parser).ParseUnverified(orderResp.OrderId, jwt.MapClaims{})
//	if err != nil {
//		return
//	}
//
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok {
//		return nil, nil, errors.InitID("can not read payment order token")
//	}
//
//	return
//}
//
//func (c *Client) TransferOTP() (err error) {
//	r, err := c.rc.R().Get("/transfer/send-otp")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		return errors.InitID(r.Status())
//	}
//
//	return
//}
//
//func (c *Client) TransferVerify(mobile, amount, otp string) (err error) {
//	//orderResp := &DepositWalletResponse{}
//
//	r, err := c.rc.R().
//		SetBody(map[string]interface{}{
//			"phoneNumber": mobile,
//			"amount":      amount,
//			"otpCode":     otp,
//		}).
//		//SetResult(depositResp).
//		Post("/transfer/verify-otp")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		return errors.InitID(r.Status())
//	}
//
//	return
//}
//
//func (c *Client) WithdrawOTP() (err error) {
//	r, err := c.rc.R().Get("/withdraw/send-otp")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		return errors.InitID(r.Status())
//	}
//
//	return
//}
//
//func (c *Client) WithdrawVerify(sheba, amount, otp string) (apiErr *APIError, err error) {
//	//orderResp := &DepositWalletResponse{}
//	apiErr = &APIError{}
//
//	r, err := c.rc.R().
//		SetBody(map[string]interface{}{
//			"sheba":   sheba,
//			"amount":  amount,
//			"otpCode": otp,
//		}).
//		SetError(apiErr).
//		//SetResult(depositResp).
//		Post("/withdraw/verify-otp")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		err = fmt.Errorf("code: %d, message: %s", apiErr.Code, apiErr.Message)
//		return
//	}
//
//	return
//}
//
//func (c *Client) TransactionsHistory(page, perPage int) (err error) {
//	//orderResp := &DepositWalletResponse{}
//
//	r, err := c.rc.R().
//		SetBody(map[string]interface{}{
//			"page":    page,
//			"perPage": perPage,
//		}).
//		//SetResult(depositResp).
//		Post("/transactions")
//	if err != nil {
//		return
//	}
//
//	if r.IsError() {
//		return errors.InitID(r.Status())
//	}
//
//	return
//}
//
//func (c *Client) RegisterAccount(nationalCode string) (resp *CreateWalletVerifyOtpResponse, err error) {
//
//	// *** check user wallet ***
//
//	// API: GetWallet
//
//	getWalletResp, err := c.GetWallet()
//	if err != nil {
//		return
//	}
//
//	// *** check wallet status ***
//	if getWalletResp.Status == "enable" {
//		return nil, errors.InitID("user has enable wallet")
//	}
//
//	// *** user create wallet ***
//
//	createWalletOtpResp, _, err := c.CreateWallet(nationalCode)
//	if err != nil {
//		return
//	}
//
//	_ = createWalletOtpResp
//
//	createWalletVerifyResp, _, err := c.CreateWalletVerify("1234")
//	if err != nil {
//		return
//	}
//
//	if createWalletVerifyResp.Status != "enable" {
//		return nil, errors.InitID("unable to create wallet, wallet is disable")
//	}
//
//	return createWalletVerifyResp, nil
//}
