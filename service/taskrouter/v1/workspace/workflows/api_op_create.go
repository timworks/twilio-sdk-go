// Package workflows contains auto-generated files. DO NOT MODIFY
package workflows

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateWorkflowInput defines the input fields for creating a new workflow resource
type CreateWorkflowInput struct {
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	Configuration                 string  `validate:"required" form:"Configuration"`
	FallbackAssignmentCallbackURL *string `form:"FallbackAssignmentCallbackUrl,omitempty"`
	FriendlyName                  string  `validate:"required" form:"FriendlyName"`
	TaskReservationTimeout        *int    `form:"TaskReservationTimeout,omitempty"`
}

// CreateWorkflowResponse defines the response fields for the created workflow
type CreateWorkflowResponse struct {
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

// Create creates a new workflow
// See https://www.twilio.com/docs/taskrouter/api/workflow#create-a-workflow-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWorkflowInput) (*CreateWorkflowResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new workflow
// See https://www.twilio.com/docs/taskrouter/api/workflow#create-a-workflow-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateWorkflowInput) (*CreateWorkflowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workflows",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateWorkflowInput{}
	}

	response := &CreateWorkflowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
