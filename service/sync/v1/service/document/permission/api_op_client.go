// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific document permissions resource
// See https://www.twilio.com/docs/sync/api/document-permission-resource for more details
type Client struct {
	client *client.Client

	documentSid string
	identity    string
	serviceSid  string
}

// ClientProperties are the properties required to manage the permission resources
type ClientProperties struct {
	DocumentSid string
	Identity    string
	ServiceSid  string
}

// New creates a new instance of the permission client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		documentSid: properties.DocumentSid,
		identity:    properties.Identity,
		serviceSid:  properties.ServiceSid,
	}
}
