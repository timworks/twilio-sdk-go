// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchChannelResponse defines the response fields for the retrieved channel
type FetchChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	TaskSid     *string    `json:"task_sid,omitempty"`
	URL         string     `json:"url"`
	UserSid     string     `json:"user_sid"`
}

// Fetch retrieves a channel resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a channel resource
func (c Client) FetchWithContext(context context.Context) (*FetchChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Channels/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
