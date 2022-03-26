// Package sessions contains auto-generated files. DO NOT MODIFY
package sessions

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing session resources
// See https://www.twilio.com/docs/proxy/api/session for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the sessions resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the sessions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
