// Package tokens contains auto-generated files. DO NOT MODIFY
package tokens

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateTokenInput defines the input fields for generating a token
type CreateTokenInput struct {
	Ttl *int `form:"Ttl,omitempty"`
}

type CreateIceServerResponse struct {
	Credential *string `json:"credential,omitempty"`
	URL        string  `json:"url"`
	URLs       string  `json:"urls"`
	Username   *string `json:"username,omitempty"`
}

// CreateTokenResponse defines the response fields for the generated token
type CreateTokenResponse struct {
	AccountSid  string                    `json:"account_sid"`
	DateCreated utils.RFC2822Time         `json:"date_created"`
	DateUpdated *utils.RFC2822Time        `json:"date_updated,omitempty"`
	IceServers  []CreateIceServerResponse `json:"ice_servers"`
	Password    string                    `json:"password"`
	Ttl         string                    `json:"ttl"`
	Username    string                    `json:"username"`
}

// Create generate an account token
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTokenInput) (*CreateTokenResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext generate an account token
func (c Client) CreateWithContext(context context.Context, input *CreateTokenInput) (*CreateTokenResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Tokens.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateTokenInput{}
	}

	response := &CreateTokenResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
