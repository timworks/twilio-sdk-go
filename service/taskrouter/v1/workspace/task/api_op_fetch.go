// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchTaskResponse defines the response fields for the retrieved task
type FetchTaskResponse struct {
	AccountSid            string      `json:"account_sid"`
	Age                   int         `json:"age"`
	AssignmentStatus      string      `json:"assignment_status"`
	Attributes            interface{} `json:"attributes"`
	DateCreated           time.Time   `json:"date_created"`
	DateUpdated           *time.Time  `json:"date_updated,omitempty"`
	Priority              *int        `json:"priority,omitempty"`
	Reason                *string     `json:"reason,omitempty"`
	Sid                   string      `json:"sid"`
	TaskChannelSid        *string     `json:"task_channel_sid,omitempty"`
	TaskChannelUniqueName *string     `json:"task_channel_unique_name,omitempty"`
	TaskQueueEnteredDate  *time.Time  `json:"task_queue_entered_date,omitempty"`
	TaskQueueFriendlyName *string     `json:"task_queue_friendly_name,omitempty"`
	TaskQueueSid          *string     `json:"task_queue_sid,omitempty"`
	Timeout               int         `json:"timeout"`
	URL                   string      `json:"url"`
	WorkflowFriendlyName  *string     `json:"workflow_friendly_name,omitempty"`
	WorkflowSid           *string     `json:"workflow_sid,omitempty"`
	WorkspaceSid          string      `json:"workspace_sid"`
}

// Fetch retrieves an task resource
// See https://www.twilio.com/docs/taskrouter/api/task#fetch-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchTaskResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an task resource
// See https://www.twilio.com/docs/taskrouter/api/task#fetch-a-task-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchTaskResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Tasks/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchTaskResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
