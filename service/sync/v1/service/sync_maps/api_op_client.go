// Package sync_maps contains auto-generated files. DO NOT MODIFY
package sync_maps

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing map resources
// See https://www.twilio.com/docs/sync/api/map-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the syncmaps resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the syncmaps client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
