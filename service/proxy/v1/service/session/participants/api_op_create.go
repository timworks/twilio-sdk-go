// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateParticipantInput defines the input fields for creating a new participant resource
type CreateParticipantInput struct {
	FriendlyName       *string `form:"FriendlyName,omitempty"`
	Identifier         string  `validate:"required" form:"Identifier"`
	ProxyIdentifier    *string `form:"ProxyIdentifier,omitempty"`
	ProxyIdentifierSid *string `form:"ProxyIdentifierSid,omitempty"`
}

// CreateParticipantResponse defines the response fields for the created participant
type CreateParticipantResponse struct {
	AccountSid         string     `json:"account_sid"`
	DateCreated        time.Time  `json:"date_created"`
	DateDeleted        *time.Time `json:"date_deleted,omitempty"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	FriendlyName       *string    `json:"friendly_name,omitempty"`
	Identifier         string     `json:"identifier"`
	ProxyIdentifier    *string    `json:"proxy_identifier,omitempty"`
	ProxyIdentifierSid *string    `json:"proxy_identifier_sid,omitempty"`
	ServiceSid         string     `json:"service_sid"`
	SessionSid         string     `json:"session_sid"`
	Sid                string     `json:"sid"`
	URL                string     `json:"url"`
}

// Create creates a new participant
// See https://www.twilio.com/docs/proxy/api/participant#create-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new participant
// See https://www.twilio.com/docs/proxy/api/participant#create-a-participant-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateParticipantInput) (*CreateParticipantResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Sessions/{sessionSid}/Participants",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sessionSid": c.sessionSid,
		},
	}

	if input == nil {
		input = &CreateParticipantInput{}
	}

	response := &CreateParticipantResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
