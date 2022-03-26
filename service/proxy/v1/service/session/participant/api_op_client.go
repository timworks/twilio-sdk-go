// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/proxy/v1/service/session/participant/message_interactions"
)

// Client for managing a specific participant resource
// See https://www.twilio.com/docs/proxy/api/participant for more details
type Client struct {
	client *client.Client

	serviceSid string
	sessionSid string
	sid        string

	MessageInteractions *message_interactions.Client
}

// ClientProperties are the properties required to manage the participant resources
type ClientProperties struct {
	ServiceSid string
	SessionSid string
	Sid        string
}

// New creates a new instance of the participant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sessionSid: properties.SessionSid,
		sid:        properties.Sid,

		MessageInteractions: message_interactions.New(client, message_interactions.ClientProperties{
			ParticipantSid: properties.Sid,
			ServiceSid:     properties.ServiceSid,
			SessionSid:     properties.SessionSid,
		}),
	}
}
