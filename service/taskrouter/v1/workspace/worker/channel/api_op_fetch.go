// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchChannelResponse defines the response fields for the retrieved worker channel
type FetchChannelResponse struct {
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

// Fetch retrieves an worker channel resource
// See https://www.twilio.com/docs/taskrouter/api/worker-channel#fetch-a-workerchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an worker channel resource
// See https://www.twilio.com/docs/taskrouter/api/worker-channel#fetch-a-workerchannel-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/{workerSid}/Channels/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workerSid":    c.workerSid,
			"sid":          c.sid,
		},
	}

	response := &FetchChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
