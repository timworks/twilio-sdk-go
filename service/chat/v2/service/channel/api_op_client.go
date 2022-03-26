// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/invite"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/invites"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/member"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/members"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/message"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/messages"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	"github.com/timworks/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
)

// Client for managing a specific channel resource
// See https://www.twilio.com/docs/chat/rest/channel-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Invite   func(string) *invite.Client
	Invites  *invites.Client
	Member   func(string) *member.Client
	Members  *members.Client
	Message  func(string) *message.Client
	Messages *messages.Client
	Webhook  func(string) *webhook.Client
	Webhooks *webhooks.Client
}

// ClientProperties are the properties required to manage the channel resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the channel client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Invite: func(inviteSid string) *invite.Client {
			return invite.New(client, invite.ClientProperties{
				ChannelSid: properties.Sid,
				ServiceSid: properties.ServiceSid,
				Sid:        inviteSid,
			})
		},
		Invites: invites.New(client, invites.ClientProperties{
			ChannelSid: properties.Sid,
			ServiceSid: properties.ServiceSid,
		}),
		Member: func(memberSid string) *member.Client {
			return member.New(client, member.ClientProperties{
				ChannelSid: properties.Sid,
				ServiceSid: properties.ServiceSid,
				Sid:        memberSid,
			})
		},
		Members: members.New(client, members.ClientProperties{
			ChannelSid: properties.Sid,
			ServiceSid: properties.ServiceSid,
		}),
		Message: func(messageSid string) *message.Client {
			return message.New(client, message.ClientProperties{
				ChannelSid: properties.Sid,
				ServiceSid: properties.ServiceSid,
				Sid:        messageSid,
			})
		},
		Messages: messages.New(client, messages.ClientProperties{
			ChannelSid: properties.Sid,
			ServiceSid: properties.ServiceSid,
		}),
		Webhook: func(webhookSid string) *webhook.Client {
			return webhook.New(client, webhook.ClientProperties{
				ChannelSid: properties.Sid,
				ServiceSid: properties.ServiceSid,
				Sid:        webhookSid,
			})
		},
		Webhooks: webhooks.New(client, webhooks.ClientProperties{
			ChannelSid: properties.Sid,
			ServiceSid: properties.ServiceSid,
		}),
	}
}
