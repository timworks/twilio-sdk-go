// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateWebhookInput defines the input fields for creating a new webhook resource
type CreateWebhookInput struct {
	Events        string  `validate:"required" form:"Events"`
	UniqueName    string  `validate:"required" form:"UniqueName"`
	WebhookMethod *string `form:"WebhookMethod,omitempty"`
	WebhookURL    string  `validate:"required" form:"WebhookUrl"`
}

// CreateWebhookResponse defines the response fields for the created webhook
type CreateWebhookResponse struct {
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	Events        string     `json:"events"`
	Sid           string     `json:"sid"`
	URL           string     `json:"url"`
	UniqueName    string     `json:"unique_name"`
	WebhookMethod string     `json:"webhook_method"`
	WebhookURL    string     `json:"webhook_url"`
}

// Create creates a new webhook
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#create-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new webhook
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#create-a-webhook-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	if input == nil {
		input = &CreateWebhookInput{}
	}

	response := &CreateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
