// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateWebhookInput defines input fields for updating a webhook resource
type UpdateWebhookInput struct {
	Filters        *[]string `form:"Filters,omitempty"`
	Method         *string   `form:"Method,omitempty"`
	PostWebhookURL *string   `form:"PostWebhookUrl,omitempty"`
	PreWebhookURL  *string   `form:"PreWebhookUrl,omitempty"`
	Target         *string   `form:"Target,omitempty"`
}

// UpdateWebhookResponse defines the response fields for the updated webhook
type UpdateWebhookResponse struct {
	AccountSid     string   `json:"account_sid"`
	Filters        []string `json:"filters"`
	Method         string   `json:"method"`
	PostWebhookURL *string  `json:"post_webhook_url,omitempty"`
	PreWebhookURL  *string  `json:"pre_webhook_url,omitempty"`
	Target         string   `json:"target"`
	URL            string   `json:"url"`
}

// Update modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-webhook-resource#update-a-conversationwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-webhook-resource#update-a-conversationwebhook-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebhookInput) (*UpdateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Conversations/Webhooks",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &UpdateWebhookInput{}
	}

	response := &UpdateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
