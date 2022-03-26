// Package plugin contains auto-generated files. DO NOT MODIFY
package plugin

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific plugin resource
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client

	configurationSid string
	sid              string
}

// ClientProperties are the properties required to manage the plugin resources
type ClientProperties struct {
	ConfigurationSid string
	Sid              string
}

// New creates a new instance of the plugin client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		configurationSid: properties.ConfigurationSid,
		sid:              properties.Sid,
	}
}
