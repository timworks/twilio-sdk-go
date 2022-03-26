// Package task_queues contains auto-generated files. DO NOT MODIFY
package task_queues

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateTaskQueueInput defines the input fields for creating a new task queue resource
type CreateTaskQueueInput struct {
	AssignmentActivitySid  *string `form:"AssignmentActivitySid,omitempty"`
	FriendlyName           string  `validate:"required" form:"FriendlyName"`
	MaxReservedWorkers     *int    `form:"MaxReservedWorkers,omitempty"`
	ReservationActivitySid *string `form:"ReservationActivitySid,omitempty"`
	TargetWorkers          *string `form:"TargetWorkers,omitempty"`
	TaskOrder              *string `form:"TaskOrder,omitempty"`
}

// CreateTaskQueueResponse defines the response fields for the created task queue
type CreateTaskQueueResponse struct {
	AccountSid              string     `json:"account_sid"`
	AssignmentActivityName  *string    `json:"assignment_activity_name,omitempty"`
	AssignmentActivitySid   *string    `json:"assignment_activity_sid,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	EventCallbackURL        *string    `json:"event_callback_url,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	MaxReservedWorkers      int        `json:"max_reserved_workers"`
	ReservationActivityName *string    `json:"reservation_activity_name,omitempty"`
	ReservationActivitySid  *string    `json:"reservation_activity_sid,omitempty"`
	Sid                     string     `json:"sid"`
	TargetWorkers           *string    `json:"target_workers,omitempty"`
	TaskOrder               string     `json:"task_order"`
	URL                     string     `json:"url"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// Create creates a new task queue
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-create for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTaskQueueInput) (*CreateTaskQueueResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task queue
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-create for more details
func (c Client) CreateWithContext(context context.Context, input *CreateTaskQueueInput) (*CreateTaskQueueResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/TaskQueues",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateTaskQueueInput{}
	}

	response := &CreateTaskQueueResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
