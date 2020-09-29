// Package webhooks contains auto-generated files. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// CreateChannelWebhookInput defines the input fields for creating a new webhook resource
type CreateChannelWebhookInput struct {
	ConfigurationFilters    *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationFlowSid    *string   `form:"Configuration.FlowSid,omitempty"`
	ConfigurationMethod     *string   `form:"Configuration.Method,omitempty"`
	ConfigurationRetryCount *int      `form:"Configuration.RetryCount,omitempty"`
	ConfigurationTriggers   *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationURL        *string   `form:"Configuration.Url,omitempty"`
	Type                    string    `validate:"required" form:"Type"`
}

type CreateChannelWebhookResponseConfiguration struct {
	Filters    *[]string `json:"filters,omitempty"`
	FlowSid    *string   `json:"flow_sid,omitempty"`
	Method     *string   `json:"method,omitempty"`
	RetryCount *int      `json:"retry_count,omitempty"`
	Triggers   *[]string `json:"triggers,omitempty"`
	URL        *string   `json:"url,omitempty"`
}

// CreateChannelWebhookResponse defines the response fields for the created webhook
type CreateChannelWebhookResponse struct {
	AccountSid    string                                    `json:"account_sid"`
	ChannelSid    string                                    `json:"channel_sid"`
	Configuration CreateChannelWebhookResponseConfiguration `json:"configuration"`
	DateCreated   time.Time                                 `json:"date_created"`
	DateUpdated   *time.Time                                `json:"date_updated,omitempty"`
	ServiceSid    string                                    `json:"service_sid"`
	Sid           string                                    `json:"sid"`
	Type          string                                    `json:"type"`
	URL           string                                    `json:"url"`
}

// Create creates a new webhook
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#create-a-channelwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelWebhookInput) (*CreateChannelWebhookResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new webhook
// See https://www.twilio.com/docs/chat/rest/channel-webhook-resource#create-a-channelwebhook-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChannelWebhookInput) (*CreateChannelWebhookResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	if input == nil {
		input = &CreateChannelWebhookInput{}
	}

	response := &CreateChannelWebhookResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
