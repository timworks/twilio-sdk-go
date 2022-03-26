// Package alpha_senders contains auto-generated files. DO NOT MODIFY
package alpha_senders

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing alpha sender resources
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the alphasenders resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the alphasenders client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
