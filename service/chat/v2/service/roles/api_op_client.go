// Package roles contains auto-generated files. DO NOT MODIFY
package roles

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing role resources
// See https://www.twilio.com/docs/chat/rest/role-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the roles resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the roles client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
