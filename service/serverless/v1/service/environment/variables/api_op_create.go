// Package variables contains auto-generated files. DO NOT MODIFY
package variables

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateVariableInput defines the input fields for creating a new environment variable resource
type CreateVariableInput struct {
	Key   string `validate:"required" form:"Key"`
	Value string `validate:"required" form:"Value"`
}

// CreateVariableResponse defines the response fields for the created environment variable
type CreateVariableResponse struct {
	AccountSid     string     `json:"account_sid"`
	DateCreated    time.Time  `json:"date_created"`
	DateUpdated    *time.Time `json:"date_updated,omitempty"`
	EnvironmentSid string     `json:"environment_sid"`
	Key            string     `json:"key"`
	ServiceSid     string     `json:"service_sid"`
	Sid            string     `json:"sid"`
	URL            string     `json:"url"`
	Value          string     `json:"value"`
}

// Create creates a new environment variable
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#create-a-variable-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateVariableInput) (*CreateVariableResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new environment variable
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#create-a-variable-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateVariableInput) (*CreateVariableResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Environments/{environmentSid}/Variables",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
		},
	}

	if input == nil {
		input = &CreateVariableInput{}
	}

	response := &CreateVariableResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
