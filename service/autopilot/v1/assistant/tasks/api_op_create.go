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
	Actions      *string `form:"Actions,omitempty"`
	ActionsURL   *string `form:"ActionsUrl,omitempty"`
	FriendlyName *string `form:"FriendlyName,omitempty"`
	UniqueName   string  `validate:"required" form:"UniqueName"`
}

// CreateTaskResponse defines the response fields for the created task
type CreateTaskResponse struct {
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

// Create creates a new task
// See https://www.twilio.com/docs/autopilot/api/task#create-a-task-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateTaskInput) (*CreateTaskResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new task
// See https://www.twilio.com/docs/autopilot/api/task#create-a-task-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateTaskInput) (*CreateTaskResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
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
