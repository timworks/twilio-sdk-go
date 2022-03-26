// Package session contains auto-generated files. DO NOT MODIFY
package session

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/proxy/v1/service/session/interaction"
	"github.com/timworks/twilio-sdk-go/service/proxy/v1/service/session/interactions"
	"github.com/timworks/twilio-sdk-go/service/proxy/v1/service/session/participant"
	"github.com/timworks/twilio-sdk-go/service/proxy/v1/service/session/participants"
)

// Client for managing a specific session resource
// See https://www.twilio.com/docs/proxy/api/session for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Interaction  func(string) *interaction.Client
	Interactions *interactions.Client
	Participant  func(string) *participant.Client
	Participants *participants.Client
}

// ClientProperties are the properties required to manage the session resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the session client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Interaction: func(interactionSid string) *interaction.Client {
			return interaction.New(client, interaction.ClientProperties{
				ServiceSid: properties.ServiceSid,
				SessionSid: properties.Sid,
				Sid:        interactionSid,
			})
		},
		Interactions: interactions.New(client, interactions.ClientProperties{
			ServiceSid: properties.ServiceSid,
			SessionSid: properties.Sid,
		}),
		Participant: func(participantSid string) *participant.Client {
			return participant.New(client, participant.ClientProperties{
				ServiceSid: properties.ServiceSid,
				SessionSid: properties.Sid,
				Sid:        participantSid,
			})
		},
		Participants: participants.New(client, participants.ClientProperties{
			ServiceSid: properties.ServiceSid,
			SessionSid: properties.Sid,
		}),
	}
}
