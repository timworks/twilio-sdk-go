// Package keys contains auto-generated files. DO NOT MODIFY
package keys

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateKeyInput defines input fields for creating a new API Key
type CreateKeyInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// CreateKeyResponse defines the response fields for creating a new API Key
type CreateKeyResponse struct {
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Secret       string             `json:"secret"`
	Sid          string             `json:"sid"`
}

// Create creates a new key resource
// See https://www.twilio.com/docs/iam/keys/api-key-resource#create-a-new-api-key for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateKeyInput) (*CreateKeyResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new key resource
// See https://www.twilio.com/docs/iam/keys/api-key-resource#create-a-new-api-key for more details
func (c Client) CreateWithContext(context context.Context, input *CreateKeyInput) (*CreateKeyResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Keys.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateKeyInput{}
	}

	response := &CreateKeyResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
