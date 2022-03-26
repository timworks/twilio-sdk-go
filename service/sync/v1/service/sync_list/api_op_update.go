// Package sync_list contains auto-generated files. DO NOT MODIFY
package sync_list

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSyncListInput defines input fields for updating a list resource
type UpdateSyncListInput struct {
	CollectionTtl *int `form:"CollectionTtl,omitempty"`
	Ttl           *int `form:"Ttl,omitempty"`
}

// UpdateSyncListResponse defines the response fields for the updated list
type UpdateSyncListResponse struct {
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

// Update modifies a list resource
// See https://www.twilio.com/docs/sync/api/list-resource#update-a-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncListInput) (*UpdateSyncListResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a list resource
// See https://www.twilio.com/docs/sync/api/list-resource#update-a-list-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncListInput) (*UpdateSyncListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateSyncListInput{}
	}

	response := &UpdateSyncListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
