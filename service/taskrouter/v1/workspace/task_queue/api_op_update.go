// Package task_queue contains auto-generated files. DO NOT MODIFY
package task_queue

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateTaskQueueInput defines input fields for updating a task queue resource
type UpdateTaskQueueInput struct {
	AssignmentActivitySid  *string `form:"AssignmentActivitySid,omitempty"`
	FriendlyName           *string `form:"FriendlyName,omitempty"`
	MaxReservedWorkers     *int    `form:"MaxReservedWorkers,omitempty"`
	ReservationActivitySid *string `form:"ReservationActivitySid,omitempty"`
	TargetWorkers          *string `form:"TargetWorkers,omitempty"`
	TaskOrder              *string `form:"TaskOrder,omitempty"`
}

// UpdateTaskQueueResponse defines the response fields for the updated task queue
type UpdateTaskQueueResponse struct {
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

// Update modifies a task queue resource
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-update for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTaskQueueInput) (*UpdateTaskQueueResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task queue resource
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-update for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTaskQueueInput) (*UpdateTaskQueueResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/TaskQueues/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateTaskQueueInput{}
	}

	response := &UpdateTaskQueueResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
