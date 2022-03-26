// Package assistant contains auto-generated files. DO NOT MODIFY
package assistant

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/defaults"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/dialogue"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/field_type"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/field_types"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/model_build"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/model_builds"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/queries"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/query"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/style_sheet"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/task"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/tasks"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/webhook"
	"github.com/timworks/twilio-sdk-go/service/autopilot/v1/assistant/webhooks"
)

// Client for managing a specific assistant resource
// See https://www.twilio.com/docs/autopilot/api/assistant for more details
type Client struct {
	client *client.Client

	sid string

	Defaults    func() *defaults.Client
	Dialogue    func(string) *dialogue.Client
	FieldType   func(string) *field_type.Client
	FieldTypes  *field_types.Client
	ModelBuild  func(string) *model_build.Client
	ModelBuilds *model_builds.Client
	Queries     *queries.Client
	Query       func(string) *query.Client
	StyleSheet  func() *style_sheet.Client
	Task        func(string) *task.Client
	Tasks       *tasks.Client
	Webhook     func(string) *webhook.Client
	Webhooks    *webhooks.Client
}

// ClientProperties are the properties required to manage the assistant resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the assistant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Defaults: func() *defaults.Client {
			return defaults.New(client, defaults.ClientProperties{
				AssistantSid: properties.Sid,
			})
		},
		Dialogue: func(dialogueSid string) *dialogue.Client {
			return dialogue.New(client, dialogue.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          dialogueSid,
			})
		},
		FieldType: func(fieldTypeSid string) *field_type.Client {
			return field_type.New(client, field_type.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          fieldTypeSid,
			})
		},
		FieldTypes: field_types.New(client, field_types.ClientProperties{
			AssistantSid: properties.Sid,
		}),
		ModelBuild: func(modelBuildSid string) *model_build.Client {
			return model_build.New(client, model_build.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          modelBuildSid,
			})
		},
		ModelBuilds: model_builds.New(client, model_builds.ClientProperties{
			AssistantSid: properties.Sid,
		}),
		Queries: queries.New(client, queries.ClientProperties{
			AssistantSid: properties.Sid,
		}),
		Query: func(querySid string) *query.Client {
			return query.New(client, query.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          querySid,
			})
		},
		StyleSheet: func() *style_sheet.Client {
			return style_sheet.New(client, style_sheet.ClientProperties{
				AssistantSid: properties.Sid,
			})
		},
		Task: func(taskSid string) *task.Client {
			return task.New(client, task.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          taskSid,
			})
		},
		Tasks: tasks.New(client, tasks.ClientProperties{
			AssistantSid: properties.Sid,
		}),
		Webhook: func(webhookSid string) *webhook.Client {
			return webhook.New(client, webhook.ClientProperties{
				AssistantSid: properties.Sid,
				Sid:          webhookSid,
			})
		},
		Webhooks: webhooks.New(client, webhooks.ClientProperties{
			AssistantSid: properties.Sid,
		}),
	}
}
