// Package variable contains auto-generated files. DO NOT MODIFY
package variable

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateVariableInput defines input fields for updating a environment variable resource
type UpdateVariableInput struct {
	Key   *string `form:"Key,omitempty"`
	Value *string `form:"Value,omitempty"`
}

// UpdateVariableResponse defines the response fields for the updated environment variable
type UpdateVariableResponse struct {
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

// Update modifies a environment variable resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#update-a-variable-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateVariableInput) (*UpdateVariableResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a environment variable resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/variable#update-a-variable-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateVariableInput) (*UpdateVariableResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Environments/{environmentSid}/Variables/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"environmentSid": c.environmentSid,
			"sid":            c.sid,
		},
	}

	if input == nil {
		input = &UpdateVariableInput{}
	}

	response := &UpdateVariableResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
