// Package account contains auto-generated files. DO NOT MODIFY
package account

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchAccountResponse defines the response fields for the retrieved account
type FetchAccountResponse struct {
	AuthToken       string             `json:"auth_token"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	OwnerAccountSid string             `json:"owner_account_sid"`
	Sid             string             `json:"sid"`
	Status          string             `json:"status"`
	Type            string             `json:"type"`
}

// Fetch retrieves a Twilio account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#fetch-an-account-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchAccountResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a Twilio account (parent or sub account) resource
// See https://www.twilio.com/docs/iam/api/account#fetch-an-account-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchAccountResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{sid}.json",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchAccountResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
