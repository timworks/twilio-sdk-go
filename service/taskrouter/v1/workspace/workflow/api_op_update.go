// Package workflow contains auto-generated files. DO NOT MODIFY
package workflow

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateWorkflowInput defines input fields for updating a workflow resource
type UpdateWorkflowInput struct {
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	Configuration                 *string `form:"Configuration,omitempty"`
	FallbackAssignmentCallbackURL *string `form:"FallbackAssignmentCallbackUrl,omitempty"`
	FriendlyName                  *string `form:"FriendlyName,omitempty"`
	ReEvaluateTasks               *bool   `form:"ReEvaluateTasks,omitempty"`
	TaskReservationTimeout        *int    `form:"TaskReservationTimeout,omitempty"`
}

// UpdateWorkflowResponse defines the response fields for the updated workflow
type UpdateWorkflowResponse struct {
	AccountSid                    string     `json:"account_sid"`
	AssignmentCallbackURL         *string    `json:"assignment_callback_url,omitempty"`
	Configuration                 string     `json:"configuration"`
	DateCreated                   time.Time  `json:"date_created"`
	DateUpdated                   *time.Time `json:"date_updated,omitempty"`
	DocumentContentType           string     `json:"document_content_type"`
	FallbackAssignmentCallbackURL *string    `json:"fallback_assignment_callback_url,omitempty"`
	FriendlyName                  string     `json:"friendly_name"`
	Sid                           string     `json:"sid"`
	TaskReservationTimeout        int        `json:"task_reservation_timeout"`
	URL                           string     `json:"url"`
	WorkspaceSid                  string     `json:"workspace_sid"`
}

// Update modifies a workflow resource
// See https://www.twilio.com/docs/taskrouter/api/workflow#update-a-workflow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWorkflowInput) (*UpdateWorkflowResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a workflow resource
// See https://www.twilio.com/docs/taskrouter/api/workflow#update-a-workflow-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWorkflowInput) (*UpdateWorkflowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workflows/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateWorkflowInput{}
	}

	response := &UpdateWorkflowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
