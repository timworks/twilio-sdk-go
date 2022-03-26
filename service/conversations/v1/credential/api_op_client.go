// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific role resource
// See https://www.twilio.com/docs/conversations/api/credential-resource for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the credential resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the credential client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
