// Package participants contains auto-generated files. DO NOT MODIFY
package participants

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing participant resources
// See https://www.twilio.com/docs/voice/api/conference-participant-resource for more details
type Client struct {
	client *client.Client

	accountSid    string
	conferenceSid string
}

// ClientProperties are the properties required to manage the participants resources
type ClientProperties struct {
	AccountSid    string
	ConferenceSid string
}

// New creates a new instance of the participants client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:    properties.AccountSid,
		conferenceSid: properties.ConferenceSid,
	}
}
