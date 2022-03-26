// Package activity contains auto-generated files. DO NOT MODIFY
package activity

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateActivityInput defines input fields for updating a activity resource
type UpdateActivityInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
}

// UpdateActivityResponse defines the response fields for the updated activity
type UpdateActivityResponse struct {
	AccountSid   string     `json:"account_sid"`
	Available    bool       `json:"available"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	WorkspaceSid string     `json:"workspace_sid"`
}

// Update modifies a activity resource
// See https://www.twilio.com/docs/taskrouter/api/activity#update-an-activity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateActivityInput) (*UpdateActivityResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a activity resource
// See https://www.twilio.com/docs/taskrouter/api/activity#update-an-activity-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateActivityInput) (*UpdateActivityResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Activities/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	if input == nil {
		input = &UpdateActivityInput{}
	}

	response := &UpdateActivityResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
