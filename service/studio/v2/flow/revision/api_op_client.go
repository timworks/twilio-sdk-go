// Package revision contains auto-generated files. DO NOT MODIFY
package revision

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific flow revision resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision for more details
type Client struct {
	client *client.Client

	flowSid        string
	revisionNumber int
}

// ClientProperties are the properties required to manage the revision resources
type ClientProperties struct {
	FlowSid        string
	RevisionNumber int
}

// New creates a new instance of the revision client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid:        properties.FlowSid,
		revisionNumber: properties.RevisionNumber,
	}
}
