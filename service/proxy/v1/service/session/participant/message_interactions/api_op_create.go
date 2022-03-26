// Package message_interactions contains auto-generated files. DO NOT MODIFY
package message_interactions

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateMessageInteractionInput defines the input fields for creating a new message interaction resource
type CreateMessageInteractionInput struct {
	Body     *string `form:"Body,omitempty"`
	MediaURL *string `form:"MediaUrl,omitempty"`
}

// CreateMessageInteractionResponse defines the response fields for the created message interaction
type CreateMessageInteractionResponse struct {
	AccountSid             string     `json:"account_sid"`
	Data                   *string    `json:"data,omitempty"`
	DateCreated            time.Time  `json:"date_created"`
	DateUpdated            *time.Time `json:"date_updated,omitempty"`
	InboundParticipantSid  *string    `json:"inbound_participant_sid,omitempty"`
	InboundResourceSid     *string    `json:"inbound_resource_sid,omitempty"`
	InboundResourceStatus  *string    `json:"inbound_resource_status,omitempty"`
	InboundResourceType    *string    `json:"inbound_resource_type,omitempty"`
	InboundResourceURL     *string    `json:"inbound_resource_url,omitempty"`
	OutboundParticipantSid *string    `json:"outbound_participant_sid,omitempty"`
	OutboundResourceSid    *string    `json:"outbound_resource_sid,omitempty"`
	OutboundResourceStatus *string    `json:"outbound_resource_status,omitempty"`
	OutboundResourceType   *string    `json:"outbound_resource_type,omitempty"`
	OutboundResourceURL    *string    `json:"outbound_resource_url,omitempty"`
	ParticipantSid         string     `json:"participant_sid"`
	ServiceSid             string     `json:"service_sid"`
	SessionSid             string     `json:"session_sid"`
	Sid                    string     `json:"sid"`
	Type                   *string    `json:"type,omitempty"`
	URL                    string     `json:"url"`
}

// Create creates a new message interaction
// See https://www.twilio.com/docs/proxy/api/sending-messages#create-a-messageinteraction-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateMessageInteractionInput) (*CreateMessageInteractionResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new message interaction
// See https://www.twilio.com/docs/proxy/api/sending-messages#create-a-messageinteraction-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateMessageInteractionInput) (*CreateMessageInteractionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Sessions/{sessionSid}/Participants/{participantSid}/MessageInteractions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":     c.serviceSid,
			"sessionSid":     c.sessionSid,
			"participantSid": c.participantSid,
		},
	}

	if input == nil {
		input = &CreateMessageInteractionInput{}
	}

	response := &CreateMessageInteractionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
