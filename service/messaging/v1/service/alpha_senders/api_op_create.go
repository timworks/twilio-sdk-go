// Package alpha_senders contains auto-generated files. DO NOT MODIFY
package alpha_senders

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateAlphaSenderInput defines the input fields for creating a new alpha sender resource
type CreateAlphaSenderInput struct {
	AlphaSender string `validate:"required" form:"AlphaSender"`
}

// CreateAlphaSenderResponse defines the response fields for the created alpha sender
type CreateAlphaSenderResponse struct {
	AccountSid   string     `json:"account_sid"`
	AlphaSender  string     `json:"alpha_sender"`
	Capabilities []string   `json:"capabilities"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Create creates a new alpha sender
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#create-an-alphasender-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateAlphaSenderInput) (*CreateAlphaSenderResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new alpha sender
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#create-an-alphasender-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateAlphaSenderInput) (*CreateAlphaSenderResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/AlphaSenders",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateAlphaSenderInput{}
	}

	response := &CreateAlphaSenderResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
