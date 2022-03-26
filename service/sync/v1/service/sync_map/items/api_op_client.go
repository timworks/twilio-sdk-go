// Package items contains auto-generated files. DO NOT MODIFY
package items

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing map item resources
// See https://www.twilio.com/docs/sync/api/map-item-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	syncMapSid string
}

// ClientProperties are the properties required to manage the items resources
type ClientProperties struct {
	ServiceSid string
	SyncMapSid string
}

// New creates a new instance of the items client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		syncMapSid: properties.SyncMapSid,
	}
}
