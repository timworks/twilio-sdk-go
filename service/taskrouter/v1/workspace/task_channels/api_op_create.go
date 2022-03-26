// Package task_channels contains auto-generated files. DO NOT MODIFY
package task_channels

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateTaskChannelInput defines the input fields for creating a new task channel resource
type CreateTaskChannelInput struct {
	ChannelOptimizedRouting *bool  `form:"ChannelOptimizedRouting,omitempty"`
	FriendlyName            string `validate:"required" form:"FriendlyName"`
	UniqueName              string `validate:"required" form:"UniqueName"`
}

// CreateTaskChannelResponse defines the response fields for the created task channel
type CreateTaskChannelResponse struct {
	AccountSid              string     `json:"account_sid"`
	ChannelOptimizedRouting *bool      `json:"channel_optimized_routing,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// Create creates a new task channel
// See https://www.twilio.com/docs/taskrouter/api/task-channel#create-a-taskchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTaskChannelInput) (*CreateTaskChannelResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task channel
// See https://www.twilio.com/docs/taskrouter/api/task-channel#create-a-taskchannel-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateTaskChannelInput) (*CreateTaskChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/TaskChannels",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateTaskChannelInput{}
	}

	response := &CreateTaskChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
