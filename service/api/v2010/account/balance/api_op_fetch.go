// Package balance contains auto-generated files. DO NOT MODIFY
package balance

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchBalanceResponse defines the response fields for retrieving an account balance
type FetchBalanceResponse struct {
	AccountSid string `json:"account_sid"`
	Balance    string `json:"balance"`
	Currency   string `json:"currency"`
}

// Fetch retrieves the balance resource for the account
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchBalanceResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the balance resource for the account
func (c Client) FetchWithContext(context context.Context) (*FetchBalanceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Balance.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	response := &FetchBalanceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
