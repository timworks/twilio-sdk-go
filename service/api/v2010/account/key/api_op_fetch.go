// Package key contains auto-generated files. DO NOT MODIFY
package key

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchKeyResponse defines the response fields for retrieving a API Key
type FetchKeyResponse struct {
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves the api key resource. The secret is not returned for security reasons
// See https://www.twilio.com/docs/iam/keys/api-key-resource#fetch-a-key-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchKeyResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the api key resource. The secret is not returned for security reasons
// See https://www.twilio.com/docs/iam/keys/api-key-resource#fetch-a-key-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchKeyResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Keys/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchKeyResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
