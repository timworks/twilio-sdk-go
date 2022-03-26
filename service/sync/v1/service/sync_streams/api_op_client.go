// Package sync_streams contains auto-generated files. DO NOT MODIFY
package sync_streams

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing stream resources
// See https://www.twilio.com/docs/sync/api/stream-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the syncstreams resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the syncstreams client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
