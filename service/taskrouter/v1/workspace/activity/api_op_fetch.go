// Package activity contains auto-generated files. DO NOT MODIFY
package activity

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchActivityResponse defines the response fields for the retrieved activity
type FetchActivityResponse struct {
	AccountSid   string     `json:"account_sid"`
	Available    bool       `json:"available"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	WorkspaceSid string     `json:"workspace_sid"`
}

// Fetch retrieves an activity resource
// See https://www.twilio.com/docs/taskrouter/api/activity#fetch-an-activity-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchActivityResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an activity resource
// See https://www.twilio.com/docs/taskrouter/api/activity#fetch-an-activity-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchActivityResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Activities/{sid}",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	response := &FetchActivityResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
