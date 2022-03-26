// Package workspace contains auto-generated files. DO NOT MODIFY
package workspace

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateWorkspaceInput defines input fields for updating a workspace resource
type UpdateWorkspaceInput struct {
	DefaultActivitySid   *string `form:"DefaultActivitySid,omitempty"`
	EventCallbackURL     *string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `form:"EventsFilter,omitempty"`
	FriendlyName         *string `form:"FriendlyName,omitempty"`
	MultiTaskEnabled     *bool   `form:"MultiTaskEnabled,omitempty"`
	PrioritizeQueueOrder *string `form:"PrioritizeQueueOrder,omitempty"`
	TimeoutActivitySid   *string `form:"TimeoutActivitySid,omitempty"`
}

// UpdateWorkspaceResponse defines the response fields for the updated workspace
type UpdateWorkspaceResponse struct {
	AccountSid           string     `json:"account_sid"`
	DateCreated          time.Time  `json:"date_created"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	DefaultActivityName  string     `json:"default_activity_name"`
	DefaultActivitySid   string     `json:"default_activity_sid"`
	EventCallbackURL     *string    `json:"event_callback_url,omitempty"`
	EventsFilter         *string    `json:"events_filter,omitempty"`
	FriendlyName         string     `json:"friendly_name"`
	MultiTaskEnabled     bool       `json:"multi_task_enabled"`
	PrioritizeQueueOrder string     `json:"prioritize_queue_order"`
	Sid                  string     `json:"sid"`
	TimeoutActivityName  string     `json:"timeout_activity_name"`
	TimeoutActivitySid   string     `json:"timeout_activity_sid"`
	URL                  string     `json:"url"`
}

// Update modifies a workspace resource
// See https://www.twilio.com/docs/taskrouter/api/workspace#update-a-workspace-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWorkspaceInput) (*UpdateWorkspaceResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a workspace resource
// See https://www.twilio.com/docs/taskrouter/api/workspace#update-a-workspace-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateWorkspaceInput) (*UpdateWorkspaceResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateWorkspaceInput{}
	}

	response := &UpdateWorkspaceResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
