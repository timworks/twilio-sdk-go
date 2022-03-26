// Package sync_stream contains auto-generated files. DO NOT MODIFY
package sync_stream

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSyncStreamInput defines input fields for updating a stream resource
type UpdateSyncStreamInput struct {
	Ttl *int `form:"Ttl,omitempty"`
}

// UpdateSyncStreamResponse defines the response fields for the updated stream
type UpdateSyncStreamResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// Update modifies an stream resource
// See https://www.twilio.com/docs/sync/api/stream-resource#update-a-sync-stream-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncStreamInput) (*UpdateSyncStreamResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an stream resource
// See https://www.twilio.com/docs/sync/api/stream-resource#update-a-sync-stream-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncStreamInput) (*UpdateSyncStreamResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Streams/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateSyncStreamInput{}
	}

	response := &UpdateSyncStreamResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
