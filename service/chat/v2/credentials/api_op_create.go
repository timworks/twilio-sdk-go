// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateCredentialInput defines the input fields for creating a new credentials resource
type CreateCredentialInput struct {
	ApiKey       *string `form:"ApiKey,omitempty"`
	Certificate  *string `form:"Certificate,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	PrivateKey   *string `form:"PrivateKey,omitempty"`
	Sandbox      *bool   `form:"Sandbox,omitempty"`
	Secret       *string `form:"Secret,omitempty"`
	Type         string  `validate:"required" form:"Type"`
}

// CreateCredentialResponse defines the response fields for the created credentials
type CreateCredentialResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sandbox      *string    `json:"sandbox,omitempty"`
	Sid          string     `json:"sid"`
	Type         string     `json:"type"`
	URL          string     `json:"url"`
}

// Create creates a new credentials
// See https://www.twilio.com/docs/chat/rest/credential-resource#create-a-credential-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateCredentialInput) (*CreateCredentialResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new credentials
// See https://www.twilio.com/docs/chat/rest/credential-resource#create-a-credential-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateCredentialInput) (*CreateCredentialResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Credentials",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateCredentialInput{}
	}

	response := &CreateCredentialResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
