// Package worker contains auto-generated files. DO NOT MODIFY
package worker

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateWorkerInput defines input fields for updating a worker resource
type UpdateWorkerInput struct {
	ActivitySid               *string `form:"ActivitySid,omitempty"`
	Attributes                *string `form:"Attributes,omitempty"`
	FriendlyName              *string `form:"FriendlyName,omitempty"`
	RejectPendingReservations *bool   `form:"RejectPendingReservations,omitempty"`
}

// UpdateWorkflowResponse defines the response fields for the updated worker
type UpdateWorkflowResponse struct {
	AccountSid        string     `json:"account_sid"`
	ActivityName      string     `json:"activity_name"`
	ActivitySid       string     `json:"activity_sid"`
	Attributes        string     `json:"attributes"`
	Available         bool       `json:"available"`
	DateCreated       time.Time  `json:"date_created"`
	DateStatusChanged *time.Time `json:"date_status_changed,omitempty"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	FriendlyName      string     `json:"friendly_name"`
	Sid               string     `json:"sid"`
	URL               string     `json:"url"`
	WorkspaceSid      string     `json:"workspace_sid"`
}

// Update modifies a worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker#update-a-worker-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWorkerInput) (*UpdateWorkflowResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a worker resource
// See https://www.twilio.com/docs/taskrouter/api/worker#update-a-worker-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWorkerInput) (*UpdateWorkflowResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workers/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateWorkerInput{}
	}

	response := &UpdateWorkflowResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
