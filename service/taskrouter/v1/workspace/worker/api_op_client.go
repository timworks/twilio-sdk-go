// Package worker contains auto-generated files. DO NOT MODIFY
package worker

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker/channel"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker/channels"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker/reservation"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker/reservations"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker/statistics"
)

// Client for managing a specific worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker for more details
type Client struct {
	client *client.Client

	sid          string
	workspaceSid string

	Channel      func(string) *channel.Client
	Channels     *channels.Client
	Reservation  func(string) *reservation.Client
	Reservations *reservations.Client
	Statistics   func() *statistics.Client
}

// ClientProperties are the properties required to manage the worker resources
type ClientProperties struct {
	Sid          string
	WorkspaceSid string
}

// New creates a new instance of the worker client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid:          properties.Sid,
		workspaceSid: properties.WorkspaceSid,

		Channel: func(channelSid string) *channel.Client {
			return channel.New(client, channel.ClientProperties{
				Sid:          channelSid,
				WorkerSid:    properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Channels: channels.New(client, channels.ClientProperties{
			WorkerSid:    properties.Sid,
			WorkspaceSid: properties.WorkspaceSid,
		}),
		Reservation: func(channelSid string) *reservation.Client {
			return reservation.New(client, reservation.ClientProperties{
				Sid:          channelSid,
				WorkerSid:    properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
		Reservations: reservations.New(client, reservations.ClientProperties{
			WorkerSid:    properties.Sid,
			WorkspaceSid: properties.WorkspaceSid,
		}),
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				WorkerSid:    properties.Sid,
				WorkspaceSid: properties.WorkspaceSid,
			})
		},
	}
}
