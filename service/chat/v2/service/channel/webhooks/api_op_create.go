// This is an autogenerated file. DO NOT MODIFY
package webhooks

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateChannelWebhookInput struct {
	Type                    string    `validate:"required" form:"Type"`
	ConfigurationUrl        *string   `form:"Configuration.Url,omitempty"`
	ConfigurationMethod     *string   `form:"Configuration.Method,omitempty"`
	ConfigurationFilters    *[]string `form:"Configuration.Filters,omitempty"`
	ConfigurationTriggers   *[]string `form:"Configuration.Triggers,omitempty"`
	ConfigurationFlowSid    *string   `form:"Configuration.FlowSid,omitempty"`
	ConfigurationRetryCount *int      `form:"Configuration.RetryCount,omitempty"`
}

type CreateChannelWebhookOutputConfiguration struct {
	Url        *string   `json:"url,omitempty"`
	Method     *string   `json:"method,omitempty"`
	Filters    *[]string `json:"filters,omitempty"`
	Triggers   *[]string `json:"triggers,omitempty"`
	FlowSid    *string   `json:"flow_sid,omitempty"`
	RetryCount *int      `json:"retry_count,omitempty"`
}

type CreateChannelWebhookOutput struct {
	Sid           string                                  `json:"sid"`
	AccountSid    string                                  `json:"account_sid"`
	ServiceSid    string                                  `json:"service_sid"`
	ChannelSid    string                                  `json:"channel_sid"`
	Type          string                                  `json:"type"`
	Configuration CreateChannelWebhookOutputConfiguration `json:"configuration"`
	DateCreated   time.Time                               `json:"date_created"`
	DateUpdated   *time.Time                              `json:"date_updated,omitempty"`
	URL           string                                  `json:"url"`
}

func (c Client) Create(input *CreateChannelWebhookInput) (*CreateChannelWebhookOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateChannelWebhookInput) (*CreateChannelWebhookOutput, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Webhooks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	output := &CreateChannelWebhookOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
