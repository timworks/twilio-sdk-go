// Package flows contains auto-generated files. DO NOT MODIFY
package flows

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateFlowInput defines the input fields for creating a new flow resource
type CreateFlowInput struct {
	CommitMessage *string `form:"CommitMessage,omitempty"`
	Definition    string  `validate:"required" form:"Definition"`
	FriendlyName  string  `validate:"required" form:"FriendlyName"`
	Status        string  `validate:"required" form:"Status"`
}

// CreateFlowResponse defines the response fields for the created flow
type CreateFlowResponse struct {
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

// Create creates a new flow
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#create-a-flow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFlowInput) (*CreateFlowResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new flow
// See https://www.twilio.com/docs/studio/rest-api/v2/flow#create-a-flow-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFlowInput) (*CreateFlowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Flows",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateFlowInput{}
	}

	response := &CreateFlowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
