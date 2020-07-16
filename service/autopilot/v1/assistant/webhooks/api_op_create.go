// This is an autogenerated file. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateWebhookInput struct {
	UniqueName    string  `validate:"required" form:"UniqueName"`
	Events        string  `validate:"required" form:"Events"`
	WebhookURL    string  `validate:"required" form:"WebhookUrl"`
	WebhookMethod *string `form:"WebhookMethod,omitempty"`
}

type CreateWebhookResponse struct {
	Sid           string     `json:"sid"`
	AccountSid    string     `json:"account_sid"`
	AssistantSid  string     `json:"assistant_sid"`
	UniqueName    string     `json:"unique_name"`
	Events        string     `json:"events"`
	WebhookURL    string     `json:"webhook_url"`
	WebhookMethod string     `json:"webhook_method"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	URL           string     `json:"url"`
}

func (c Client) Create(input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateWebhookInput) (*CreateWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &CreateWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
