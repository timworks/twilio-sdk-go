// Package execution contains auto-generated files. DO NOT MODIFY
package execution

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateExecutionInput defines input fields for updating a execution resource
type UpdateExecutionInput struct {
	Status string `validate:"required" form:"Status"`
}

// UpdateExecutionResponse defines the response fields for the updated execution
type UpdateExecutionResponse struct {
	AccountSid            string      `json:"account_sid"`
	ContactChannelAddress string      `json:"contact_channel_address"`
	Context               interface{} `json:"context"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	FlowSid               string      `json:"flow_sid"`
	Sid                   string      `json:"sid"`
	Status                string      `json:"status"`
	URL                   string      `json:"url"`
}

// Update modifies a execution resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#update-an-execution for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateExecutionInput) (*UpdateExecutionResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a execution resource
// See https://www.twilio.com/docs/studio/rest-api/v2/execution#update-an-execution for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateExecutionInput) (*UpdateExecutionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows/{flowSid}/Executions/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"flowSid": c.flowSid,
			"sid":     c.sid,
		},
	}

	if input == nil {
		input = &UpdateExecutionInput{}
	}

	response := &UpdateExecutionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
