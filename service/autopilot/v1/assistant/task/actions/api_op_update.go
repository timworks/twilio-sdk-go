// Package actions contains auto-generated files. DO NOT MODIFY
package actions

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateActionInput defines the input fields for updating a task action
type UpdateActionInput struct {
	Actions *string `form:"Actions,omitempty"`
}

// UpdateActionResponse defines the response fields for the updated task action
type UpdateActionResponse struct {
	AccountSid   string                 `json:"account_sid"`
	AssistantSid string                 `json:"assistant_sid"`
	Data         map[string]interface{} `json:"data"`
	TaskSid      string                 `json:"task_sid"`
	URL          string                 `json:"url"`
}

// Update modifies a task action resource
// See https://www.twilio.com/docs/autopilot/api/task-action#update-a-taskactions-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateActionInput) (*UpdateActionResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a task action resource
// See https://www.twilio.com/docs/autopilot/api/task-action#update-a-taskactions-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateActionInput) (*UpdateActionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Tasks/{taskSid}/Actions",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"taskSid":      c.taskSid,
		},
	}

	if input == nil {
		input = &UpdateActionInput{}
	}

	response := &UpdateActionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
