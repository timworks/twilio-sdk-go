// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific webhook resource
// See https://www.twilio.com/docs/verify/api/webhooks for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the webhook resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the webhook client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
