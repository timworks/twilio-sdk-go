// Package sync_lists contains auto-generated files. DO NOT MODIFY
package sync_lists

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing list resources
// See https://www.twilio.com/docs/sync/api/list-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
}

// ClientProperties are the properties required to manage the synclists resources
type ClientProperties struct {
	ServiceSid string
}

// New creates a new instance of the synclists client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
	}
}
