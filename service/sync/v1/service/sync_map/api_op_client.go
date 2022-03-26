// Package sync_map contains auto-generated files. DO NOT MODIFY
package sync_map

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/sync_map/item"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/sync_map/items"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/sync_map/permission"
	"github.com/timworks/twilio-sdk-go/service/sync/v1/service/sync_map/permissions"
)

// Client for managing a specific map resource
// See https://www.twilio.com/docs/sync/api/map-resource for more details
type Client struct {
	client *client.Client

	serviceSid string
	sid        string

	Item        func(string) *item.Client
	Items       *items.Client
	Permission  func(string) *permission.Client
	Permissions *permissions.Client
}

// ClientProperties are the properties required to manage the syncmap resources
type ClientProperties struct {
	ServiceSid string
	Sid        string
}

// New creates a new instance of the syncmap client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid: properties.ServiceSid,
		sid:        properties.Sid,

		Item: func(key string) *item.Client {
			return item.New(client, item.ClientProperties{
				Key:        key,
				ServiceSid: properties.ServiceSid,
				SyncMapSid: properties.Sid,
			})
		},
		Items: items.New(client, items.ClientProperties{
			ServiceSid: properties.ServiceSid,
			SyncMapSid: properties.Sid,
		}),
		Permission: func(identity string) *permission.Client {
			return permission.New(client, permission.ClientProperties{
				Identity:   identity,
				ServiceSid: properties.ServiceSid,
				SyncMapSid: properties.Sid,
			})
		},
		Permissions: permissions.New(client, permissions.ClientProperties{
			ServiceSid: properties.ServiceSid,
			SyncMapSid: properties.Sid,
		}),
	}
}
