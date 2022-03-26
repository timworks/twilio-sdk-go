// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateChannelInput defines input fields for updating a worker channel resource
type UpdateChannelInput struct {
	Available *bool `form:"Available,omitempty"`
	Capacity  *int  `form:"Capacity,omitempty"`
}

// UpdateChannelResponse defines the response fields for the updated worker channel
type UpdateChannelResponse struct {
	AccountSid                  string     `json:"account_sid"`
	AssignedTasks               int        `json:"assigned_tasks"`
	Available                   bool       `json:"available"`
	AvailableCapacityPercentage int        `json:"available_capacity_percentage"`
	ConfiguredCapacity          int        `json:"configured_capacity"`
	DateCreated                 time.Time  `json:"date_created"`
	DateUpdated                 *time.Time `json:"date_updated,omitempty"`
	Sid                         string     `json:"sid"`
	TaskChannelSid              string     `json:"task_channel_sid"`
	TaskChannelUniqueName       string     `json:"task_channel_unique_name"`
	URL                         string     `json:"url"`
	WorkerSid                   string     `json:"worker_sid"`
	WorkspaceSid                string     `json:"workspace_sid"`
}

// Update modifies a worker channel resource
// See https://www.twilio.com/docs/taskrouter/api/worker-channel#update-a-workerchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateChannelInput) (*UpdateChannelResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a worker channel resource
// See https://www.twilio.com/docs/taskrouter/api/worker-channel#update-a-workerchannel-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateChannelInput) (*UpdateChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workers/{workerSid}/Channels/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workerSid":    c.workerSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateChannelInput{}
	}

	response := &UpdateChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
