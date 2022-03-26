// Package item contains auto-generated files. DO NOT MODIFY
package item

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific map item resource
// See https://www.twilio.com/docs/sync/api/map-item-resource for more details
type Client struct {
	client *client.Client

	key        string
	serviceSid string
	syncMapSid string
}

// ClientProperties are the properties required to manage the item resources
type ClientProperties struct {
	Key        string
	ServiceSid string
	SyncMapSid string
}

// New creates a new instance of the item client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		key:        properties.Key,
		serviceSid: properties.ServiceSid,
		syncMapSid: properties.SyncMapSid,
	}
}
