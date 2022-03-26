// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific participants resource
// See https://www.twilio.com/docs/voice/api/conference-participant-resource for more details
type Client struct {
	client *client.Client

	accountSid    string
	conferenceSid string
	sid           string
}

// ClientProperties are the properties required to manage the participant resources
type ClientProperties struct {
	AccountSid    string
	ConferenceSid string
	Sid           string
}

// New creates a new instance of the participant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:    properties.AccountSid,
		conferenceSid: properties.ConferenceSid,
		sid:           properties.Sid,
	}
}
