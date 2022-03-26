// Package deployments contains auto-generated files. DO NOT MODIFY
package deployments

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing deployment resources
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment for more details
type Client struct {
	client *client.Client

	environmentSid string
	serviceSid     string
}

// ClientProperties are the properties required to manage the deployments resources
type ClientProperties struct {
	EnvironmentSid string
	ServiceSid     string
}

// New creates a new instance of the deployments client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		environmentSid: properties.EnvironmentSid,
		serviceSid:     properties.ServiceSid,
	}
}
