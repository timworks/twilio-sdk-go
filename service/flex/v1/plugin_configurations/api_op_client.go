// Package plugin_configurations contains auto-generated files. DO NOT MODIFY
package plugin_configurations

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing plugin configuration resources
// See https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration for more details
// This client is currently in beta and subject to change. Please use with caution
type Client struct {
	client *client.Client
}

// New creates a new instance of the plugin configurations client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
