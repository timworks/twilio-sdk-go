// Package web_channel contains auto-generated files. DO NOT MODIFY
package web_channel

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific web channel resource
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the web channel resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the web channel client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}
