// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific channel resource
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the channel resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the channel client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
