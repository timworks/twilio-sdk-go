// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing list permission resources
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource for more details
type Client struct {
	client *client.Client

	identity    string
	serviceSid  string
	syncListSid string
}

// ClientProperties are the properties required to manage the permission resources
type ClientProperties struct {
	Identity    string
	ServiceSid  string
	SyncListSid string
}

// New creates a new instance of the permission client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		identity:    properties.Identity,
		serviceSid:  properties.ServiceSid,
		syncListSid: properties.SyncListSid,
	}
}
