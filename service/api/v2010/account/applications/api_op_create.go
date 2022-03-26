// Package applications contains auto-generated files. DO NOT MODIFY
package applications

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateApplicationInput defines input fields for creating a new application
type CreateApplicationInput struct {
	APIVersion            *string `form:"ApiVersion,omitempty"`
	FriendlyName          *string `form:"FriendlyName,omitempty"`
	MessageStatusCallback *string `form:"MessageStatusCallback,omitempty"`
	SmsFallbackMethod     *string `form:"SmsFallbackMethod,omitempty"`
	SmsFallbackURL        *string `form:"SmsFallbackUrl,omitempty"`
	SmsMethod             *string `form:"SmsMethod,omitempty"`
	SmsStatusCallback     *string `form:"SmsStatusCallback,omitempty"`
	SmsURL                *string `form:"SmsUrl,omitempty"`
	StatusCallback        *string `form:"StatusCallback,omitempty"`
	StatusCallbackMethod  *string `form:"StatusCallbackMethod,omitempty"`
	VoiceCallerIDLookup   *bool   `form:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod   *string `form:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackURL      *string `form:"VoiceFallbackUrl,omitempty"`
	VoiceMethod           *string `form:"VoiceMethod,omitempty"`
	VoiceURL              *string `form:"VoiceUrl,omitempty"`
}

// CreateApplicationResponse defines the response fields for creating a new application
type CreateApplicationResponse struct {
	APIVersion            string             `json:"api_version"`
	AccountSid            string             `json:"account_sid"`
	DateCreated           utils.RFC2822Time  `json:"date_created"`
	DateUpdated           *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName          *string            `json:"friendly_name,omitempty"`
	MessageStatusCallback *string            `json:"message_status_callback,omitempty"`
	Sid                   string             `json:"sid"`
	SmsFallbackMethod     string             `json:"sms_fallback_method"`
	SmsFallbackURL        *string            `json:"sms_fallback_url,omitempty"`
	SmsMethod             string             `json:"sms_method"`
	SmsStatusCallback     *string            `json:"sms_status_callback,omitempty"`
	SmsURL                *string            `json:"sms_url,omitempty"`
	StatusCallback        *string            `json:"status_callback,omitempty"`
	StatusCallbackMethod  string             `json:"status_callback_method"`
	VoiceCallerIDLookup   bool               `json:"voice_caller_id_lookup"`
	VoiceFallbackMethod   string             `json:"voice_fallback_method"`
	VoiceFallbackURL      *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod           string             `json:"voice_method"`
	VoiceURL              *string            `json:"voice_url,omitempty"`
}

// Create creates a new application resource
// See https://www.twilio.com/docs/usage/api/applications#create-an-application-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateApplicationInput) (*CreateApplicationResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new application resource
// See https://www.twilio.com/docs/usage/api/applications#create-an-application-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateApplicationInput) (*CreateApplicationResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Applications.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	response := &CreateApplicationResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
