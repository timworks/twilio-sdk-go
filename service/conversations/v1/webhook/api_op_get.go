// This is an autogenerated file. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetWebhookOutput struct {
	AccountSid     string   `json:"account_sid"`
	Method         string   `json:"method"`
	Target         string   `json:"target"`
	Filters        []string `json:"filters"`
	PreWebhookUrl  *string  `json:"pre_webhook_url,omitempty"`
	PostWebhookUrl *string  `json:"post_webhook_url,omitempty"`
	URL            string   `json:"url"`
}

func (c Client) Get() (*GetWebhookOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetWebhookOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Conversations/Webhooks",
	}

	output := &GetWebhookOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
