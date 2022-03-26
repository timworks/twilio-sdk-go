// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific recording resource
// See https://www.twilio.com/docs/voice/api/recording for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the recording resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the recording client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}
