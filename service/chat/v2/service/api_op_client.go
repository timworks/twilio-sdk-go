// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/binding"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/bindings"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channels"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/role"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/roles"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/user"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/users"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/chat/rest/service-resource for more details
type Client struct {
	client *client.Client

	sid string

	Binding  func(string) *binding.Client
	Bindings *bindings.Client
	Channel  func(string) *channel.Client
	Channels *channels.Client
	Role     func(string) *role.Client
	Roles    *roles.Client
	User     func(string) *user.Client
	Users    *users.Client
}

// ClientProperties are the properties required to manage the service resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the service client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Binding: func(bindingSid string) *binding.Client {
			return binding.New(client, binding.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        bindingSid,
			})
		},
		Bindings: bindings.New(client, bindings.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Channel: func(channelSid string) *channel.Client {
			return channel.New(client, channel.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        channelSid,
			})
		},
		Channels: channels.New(client, channels.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Role: func(roleSid string) *role.Client {
			return role.New(client, role.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        roleSid,
			})
		},
		Roles: roles.New(client, roles.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		User: func(userSid string) *user.Client {
			return user.New(client, user.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        userSid,
			})
		},
		Users: users.New(client, users.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}
