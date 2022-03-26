// Package media_files contains auto-generated files. DO NOT MODIFY
package media_files

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing media resources
// See https://www.twilio.com/docs/fax/api/fax-media-resource for more details
type Client struct {
	client *client.Client

	faxSid string
}

// ClientProperties are the properties required to manage the media files resources
type ClientProperties struct {
	FaxSid string
}

// New creates a new instance of the media files client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		faxSid: properties.FaxSid,
	}
}
