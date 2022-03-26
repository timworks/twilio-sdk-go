// Package sync_map contains auto-generated files. DO NOT MODIFY
package sync_map

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchSyncMapResponse defines the response fields for the retrieved map
type FetchSyncMapResponse struct {
	AccountSid  string     `json:"account_sid"`
	CreatedBy   string     `json:"created_by"`
	DateCreated time.Time  `json:"date_created"`
	DateExpires *time.Time `json:"date_expires,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Revision    string     `json:"revision"`
	ServiceSid  string     `json:"service_Sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
	UniqueName  *string    `json:"unique_name,omitempty"`
}

// Fetch retrieves an map resource
// See https://www.twilio.com/docs/sync/api/map-resource#fetch-a-syncmap-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncMapResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an map resource
// See https://www.twilio.com/docs/sync/api/map-resource#fetch-a-syncmap-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncMapResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchSyncMapResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
