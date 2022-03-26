// Package users contains auto-generated files. DO NOT MODIFY
package users

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing user resources
// See https://www.twilio.com/docs/chat/rest/user-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the users resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the users client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
