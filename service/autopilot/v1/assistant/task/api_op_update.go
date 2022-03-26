// Package task contains auto-generated files. DO NOT MODIFY
package task

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateTaskInput defines the input fields for updating a task
type UpdateTaskInput struct {
	Actions      *string `form:"Actions,omitempty"`
	ActionsURL   *string `form:"ActionsUrl,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   *string `form:"UniqueName,omitempty"`
}

// UpdateTaskResponse defines the response fields for the updated task
type UpdateTaskResponse struct {
	AccountSid   string     `json:"account_sid"`
	ActionsURL   string     `json:"actions_url"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// Update modifies a task resource
// See https://www.twilio.com/docs/autopilot/api/task#update-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateTaskInput) (*UpdateTaskResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task resource
// See https://www.twilio.com/docs/autopilot/api/task#update-a-task-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateTaskInput) (*UpdateTaskResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateTaskInput{}
	}

	response := &UpdateTaskResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
