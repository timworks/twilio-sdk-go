// Package call contains auto-generated files. DO NOT MODIFY
package call

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/feedback"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/feedbacks"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/recording"
	"github.com/timworks/twilio-sdk-go/service/api/v2010/account/call/recordings"
)

// Client for managing a specific call resource
// See https://www.twilio.com/docs/voice/api/call-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Feedback   func() *feedback.Client
	Feedbacks  *feedbacks.Client
	Recording  func(string) *recording.Client
	Recordings *recordings.Client
}

// ClientProperties are the properties required to manage the call resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the call client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Feedback: func() *feedback.Client {
			return feedback.New(client, feedback.ClientProperties{
				AccountSid: properties.AccountSid,
				CallSid:    properties.Sid,
			})
		},
		Feedbacks: feedbacks.New(client, feedbacks.ClientProperties{
			AccountSid: properties.AccountSid,
			CallSid:    properties.Sid,
		}),
		Recording: func(recordingSid string) *recording.Client {
			return recording.New(client, recording.ClientProperties{
				AccountSid: properties.AccountSid,
				CallSid:    properties.Sid,
				Sid:        recordingSid,
			})
		},
		Recordings: recordings.New(client, recordings.ClientProperties{
			AccountSid: properties.AccountSid,
			CallSid:    properties.Sid,
		}),
	}
}
