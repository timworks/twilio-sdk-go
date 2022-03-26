// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchWebhookResponse defines the response fields for the retrieved webhook
type FetchWebhookResponse struct {
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

// Fetch retrieves a webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#fetch-a-webhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchWebhookResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a webhook resource
// See https://www.twilio.com/docs/autopilot/api/event-webhooks#fetch-a-webhook-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchWebhookResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	response := &FetchWebhookResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
