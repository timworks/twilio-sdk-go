// Package domains contains auto-generated files. DO NOT MODIFY
package domains

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// CreateDomainInput defines input fields for creating a new SIP domain
type CreateDomainInput struct {
	ByocTrunkSid              *string `form:"ByocTrunkSid,omitempty"`
	DomainName                string  `validate:"required" form:"DomainName"`
	EmergencyCallerSid        *string `form:"EmergencyCallerSid,omitempty"`
	EmergencyCallingEnabled   *bool   `form:"EmergencyCallingEnabled,omitempty"`
	FriendlyName              *string `form:"FriendlyName,omitempty"`
	Secure                    *bool   `form:"Secure,omitempty"`
	SipRegistration           *bool   `form:"SipRegistration,omitempty"`
	VoiceFallbackMethod       *string `form:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackURL          *string `form:"VoiceFallbackUrl,omitempty"`
	VoiceMethod               *string `form:"VoiceMethod,omitempty"`
	VoiceStatusCallbackMethod *string `form:"VoiceStatusCallbackMethod,omitempty"`
	VoiceStatusCallbackURL    *string `form:"VoiceStatusCallbackUrl,omitempty"`
	VoiceURL                  *string `form:"VoiceUrl,omitempty"`
}

// CreateDomainResponse defines the response fields for creating a new SIP domain
type CreateDomainResponse struct {
	AccountSid                string             `json:"account_sid"`
	ApiVersion                string             `json:"api_version"`
	AuthType                  *string            `json:"auth_type,omitempty"`
	ByocTrunkSid              *string            `json:"byoc_trunk_sid,omitempty"`
	DateCreated               utils.RFC2822Time  `json:"date_created"`
	DateUpdated               *utils.RFC2822Time `json:"date_updated,omitempty"`
	DomainName                string             `json:"domain_name"`
	EmergencyCallerSid        *string            `json:"emergency_caller_sid,omitempty"`
	EmergencyCallingEnabled   bool               `json:"emergency_calling_enabled"`
	FriendlyName              *string            `json:"friendly_name,omitempty"`
	Secure                    bool               `json:"secure"`
	Sid                       string             `json:"sid"`
	SipRegistration           bool               `json:"sip_registration"`
	VoiceFallbackMethod       *string            `json:"voice_fallback_method,omitempty"`
	VoiceFallbackURL          *string            `json:"voice_fallback_url,omitempty"`
	VoiceMethod               *string            `json:"voice_method,omitempty"`
	VoiceStatusCallbackMethod *string            `json:"voice_status_callback_method,omitempty"`
	VoiceStatusCallbackURL    *string            `json:"voice_status_callback_url,omitempty"`
	VoiceURL                  *string            `json:"voice_url,omitempty"`
}

// Create creates a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#create-a-sipdomain-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateDomainInput) (*CreateDomainResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a SIP domain resource
// See https://www.twilio.com/docs/voice/sip/api/sip-domain-resource#create-a-sipdomain-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateDomainInput) (*CreateDomainResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/SIP/Domains.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
	}

	if input == nil {
		input = &CreateDomainInput{}
	}

	response := &CreateDomainResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
