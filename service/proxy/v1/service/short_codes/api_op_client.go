// Package short_codes contains auto-generated files. DO NOT MODIFY
package short_codes

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing short code resources
// See https://www.twilio.com/docs/proxy/api/short-code for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the short codes resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the short codes client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
