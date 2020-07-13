// This is an autogenerated file. DO NOT MODIFY
package workspace

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateWorkspaceInput struct {
	FriendlyName         *string `form:"FriendlyName,omitempty"`
	EventCallbackUrl     *string `form:"EventCallbackUrl,omitempty"`
	EventsFilter         *string `form:"EventsFilter,omitempty"`
	MultiTaskEnabled     *bool   `form:"MultiTaskEnabled,omitempty"`
	Template             *string `form:"Template,omitempty"`
	PrioritizeQueueOrder *string `form:"PrioritizeQueueOrder,omitempty"`
}

type UpdateWorkspaceOutput struct {
	Sid                  string     `json:"sid"`
	AccountSid           string     `json:"account_sid"`
	FriendlyName         string     `json:"friendly_name"`
	EventCallbackURL     *string    `json:"event_callback_url,omitempty"`
	EventsFilter         *string    `json:"events_filter,omitempty"`
	DefaultActivityName  string     `json:"default_activity_name"`
	DefaultActivitySid   string     `json:"default_activity_sid"`
	MultiTaskEnabled     bool       `json:"multi_task_enabled"`
	PrioritizeQueueOrder string     `json:"prioritize_queue_order"`
	TimeoutActivityName  string     `json:"timeout_activity_name"`
	TimeoutActivitySid   string     `json:"timeout_activity_sid"`
	DateCreated          time.Time  `json:"date_created"`
	DateUpdated          *time.Time `json:"date_updated,omitempty"`
	URL                  string     `json:"url"`
}

func (c Client) Update(input *UpdateWorkspaceInput) (*UpdateWorkspaceOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateWorkspaceInput) (*UpdateWorkspaceOutput, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	output := &UpdateWorkspaceOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
