// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import "github.com/timworks/twilio-sdk-go/client"

// Client for managing a specific phone number resource
// See https://www.twilio.com/docs/proxy/api/phone-number for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string
}

// ClientProperties are the properties required to manage the phone number resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the phone number client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,
	}
}
