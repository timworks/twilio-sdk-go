// Package tasks contains auto-generated files. DO NOT MODIFY
package tasks

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateTaskInput defines the input fields for creating a new task resource
type CreateTaskInput struct {
	Attributes  *string `form:"Attributes,omitempty"`
	Priority    *int    `form:"Priority,omitempty"`
	TaskChannel *string `form:"TaskChannel,omitempty"`
	Timeout     *int    `form:"Timeout,omitempty"`
	WorkflowSid *string `form:"WorkflowSid,omitempty"`
}

// CreateTaskResponse defines the response fields for the created task
type CreateTaskResponse struct {
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

// Create creates a new task
// See https://www.twilio.com/docs/taskrouter/api/task#create-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTaskInput) (*CreateTaskResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task
// See https://www.twilio.com/docs/taskrouter/api/task#create-a-task-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateTaskInput) (*CreateTaskResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Tasks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateTaskInput{}
	}

	response := &CreateTaskResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
