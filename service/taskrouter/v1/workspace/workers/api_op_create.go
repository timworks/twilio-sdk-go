// Package workers contains auto-generated files. DO NOT MODIFY
package workers

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateWorkerInput defines the input fields for creating a new worker resource
type CreateWorkerInput struct {
	ActivitySid  *string `form:"ActivitySid,omitempty"`
	Attributes   *string `form:"Attributes,omitempty"`
	FriendlyName string  `validate:"required" form:"FriendlyName"`
}

// CreateWorkerResponse defines the response fields for the created worker
type CreateWorkerResponse struct {
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

// Create creates a new worker
// See https://www.twilio.com/docs/taskrouter/api/worker#create-a-worker-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateWorkerInput) (*CreateWorkerResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new worker
// See https://www.twilio.com/docs/taskrouter/api/worker#create-a-worker-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateWorkerInput) (*CreateWorkerResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Workspaces/{workspaceSid}/Workers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
	}

	if input == nil {
		input = &CreateWorkerInput{}
	}

	response := &CreateWorkerResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
