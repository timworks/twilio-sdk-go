// Package sync_map contains auto-generated files. DO NOT MODIFY
package sync_map

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSyncMapInput defines input fields for updating a map resource
type UpdateSyncMapInput struct {
	CollectionTtl *int `form:"CollectionTtl,omitempty"`
	Ttl           *int `form:"Ttl,omitempty"`
}

// UpdateSyncMapResponse defines the response fields for the updated map
type UpdateSyncMapResponse struct {
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

// Update modifies an map resource
// See https://www.twilio.com/docs/sync/api/map-resource#update-a-syncmap-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncMapInput) (*UpdateSyncMapResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an map resource
// See https://www.twilio.com/docs/sync/api/map-resource#update-a-syncmap-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncMapInput) (*UpdateSyncMapResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Maps/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateSyncMapInput{}
	}

	response := &UpdateSyncMapResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
