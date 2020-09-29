// Package flex_flow contains auto-generated files. DO NOT MODIFY
package flex_flow

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific flex flow resource
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the flex flow resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the flex flow client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
