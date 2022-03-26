// Package versions contains auto-generated files. DO NOT MODIFY
package versions

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing asset version resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset for more details
type Client struct {
	client *client.Client

	assetSid   string
	serviceSid string
}

// ClientProperties are the properties required to manage the versions resources
type ClientProperties struct {
	AssetSid   string
	ServiceSid string
}

// New creates a new instance of the versions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assetSid:   properties.AssetSid,
		serviceSid: properties.ServiceSid,
	}
}
