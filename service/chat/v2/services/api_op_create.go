// This is an autogenerated file. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateServiceInput struct {
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

type CreateServiceResponse struct {
	Sid                          string                 `json:"sid"`
	AccountSid                   string                 `json:"account_sid"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval"`
	DefaultChannelCreatorRoleSid string                 `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                 `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                 `json:"default_service_role_sid"`
	FriendlyName                 string                 `json:"friendly_name"`
	Limits                       map[string]interface{} `json:"limits"`
	Media                        map[string]interface{} `json:"media"`
	Notifications                map[string]interface{} `json:"notifications"`
	PostWebhookRetryCount        *int                   `json:"post_webhook_retry_count,omitempty"`
	PostWebhookUrl               *string                `json:"post_webhook_url,omitempty"`
	PreWebhookRetryCount         *int                   `json:"pre_webhook_retry_count,omitempty"`
	PreWebhookUrl                *string                `json:"pre_webhook_url,omitempty"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout"`
	WebhookFilters               *[]string              `json:"webhook_filters,omitempty"`
	WebhookMethod                *string                `json:"webhook_method,omitempty"`
	DateCreated                  time.Time              `json:"date_created"`
	DateUpdated                  *time.Time             `json:"date_updated,omitempty"`
	URL                          string                 `json:"url"`
}

func (c Client) Create(input *CreateServiceInput) (*CreateServiceResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services",
		ContentType: client.URLEncoded,
	}

	response := &CreateServiceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
