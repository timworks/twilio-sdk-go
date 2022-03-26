// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/asset"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/assets"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/build"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/builds"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/environment"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/environments"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/function"
	"github.com/timworks/twilio-sdk-go/service/serverless/v1/service/functions"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/service for more details
type Client struct {
	client *client.Client

	sid string

	Asset        func(string) *asset.Client
	Assets       *assets.Client
	Build        func(string) *build.Client
	Builds       *builds.Client
	Environment  func(string) *environment.Client
	Environments *environments.Client
	Function     func(string) *function.Client
	Functions    *functions.Client
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

		Asset: func(assetSid string) *asset.Client {
			return asset.New(client, asset.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        assetSid,
			})
		},
		Assets: assets.New(client, assets.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Build: func(buildSid string) *build.Client {
			return build.New(client, build.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        buildSid,
			})
		},
		Builds: builds.New(client, builds.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Environment: func(environmentSid string) *environment.Client {
			return environment.New(client, environment.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        environmentSid,
			})
		},
		Environments: environments.New(client, environments.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Function: func(functionSid string) *function.Client {
			return function.New(client, function.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        functionSid,
			})
		},
		Functions: functions.New(client, functions.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}
