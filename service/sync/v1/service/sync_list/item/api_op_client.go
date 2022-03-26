// Package item contains auto-generated files. DO NOT MODIFY
package item

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific list item resource
// See https://www.twilio.com/docs/sync/api/listitem-resource for more details
type Client struct {
	client *client.Client

	index       int
	serviceSid  string
	syncListSid string
}

// ClientProperties are the properties required to manage the item resources
type ClientProperties struct {
	Index       int
	ServiceSid  string
	SyncListSid string
}

// New creates a new instance of the item client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		index:       properties.Index,
		serviceSid:  properties.ServiceSid,
		syncListSid: properties.SyncListSid,
	}
}
