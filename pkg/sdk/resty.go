package sdk

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type Client struct {
	rc *resty.Client
}

func New(baseURL string, debug bool) *Client {
	return &Client{
		rc: resty.New().
			SetDebug(debug).
			SetBaseURL(baseURL).
			SetRetryCount(3).
			SetRetryWaitTime(5 * time.Second).
			SetRetryMaxWaitTime(20 * time.Second).
			AddRetryCondition(
				func(r *resty.Response, err error) bool {
					return r.StatusCode() == http.StatusTooManyRequests
				}),
	}
}
