// Package defaults contains auto-generated files. DO NOT MODIFY
package defaults

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateDefaultInput defines the input fields for updating defaults
type UpdateDefaultInput struct {
	Defaults *string `form:"Defaults,omitempty"`
}

// UpdateDefaultResponse defines the response fields for the updated defaults
type UpdateDefaultResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	URL          string                 `json:"url"`
}

// Update modifies a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#update-a-defaults-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a defaults resource
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults#update-a-defaults-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Defaults",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &UpdateDefaultInput{}
	}

	response := &UpdateDefaultResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
