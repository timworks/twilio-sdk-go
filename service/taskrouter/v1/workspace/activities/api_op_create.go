// Package activities contains auto-generated files. DO NOT MODIFY
package activities

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateActivityInput defines the input fields for creating a new activity resource
type CreateActivityInput struct {
	Available    *bool  `form:"Available,omitempty"`
	FriendlyName string `validate:"required" form:"FriendlyName"`
}

// CreateActivityResponse defines the response fields for the created activity
type CreateActivityResponse struct {
	AccountSid   string     `json:"account_sid"`
	Available    bool       `json:"available"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	WorkspaceSid string     `json:"workspace_sid"`
}

// Create creates a new activity
// See https://www.twilio.com/docs/taskrouter/api/activity#create-an-activity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateActivityInput) (*CreateActivityResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new activity
// See https://www.twilio.com/docs/taskrouter/api/activity#create-an-activity-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateActivityInput) (*CreateActivityResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Activities",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateActivityInput{}
	}

	response := &CreateActivityResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
