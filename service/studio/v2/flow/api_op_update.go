// Package flow contains auto-generated files. DO NOT MODIFY
package flow

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateFlowInput defines input fields for updating a flow resource
type UpdateFlowInput struct {
	CommitMessage *string `form:"CommitMessage,omitempty"`
	Definition    *string `form:"Definition,omitempty"`
	FriendlyName  *string `form:"FriendlyName,omitempty"`
	Status        string  `validate:"required" form:"Status"`
}

// UpdateFlowResponse defines the response fields for the updated flow
type UpdateFlowResponse struct {
	AccountSid    string                 `json:"account_sid"`
	CommitMessage *string                `json:"commit_message,omitempty"`
	DateCreated   time.Time              `json:"date_created"`
	DateUpdated   *time.Time             `json:"date_updated,omitempty"`
	Definition    map[string]interface{} `json:"definition"`
	Errors        *[]interface{}         `json:"errors,omitempty"`
	FriendlyName  string                 `json:"friendly_name"`
	Revision      int                    `json:"revision"`
	Sid           string                 `json:"sid"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Valid         bool                   `json:"valid"`
	Warnings      *[]interface{}         `json:"warnings,omitempty"`
	WebhookURL    string                 `json:"webhook_url"`
}

// Update modifies a flow resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#update-a-flow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateFlowInput) (*UpdateFlowResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a flow resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#update-a-flow-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateFlowInput) (*UpdateFlowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateFlowInput{}
	}

	response := &UpdateFlowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
