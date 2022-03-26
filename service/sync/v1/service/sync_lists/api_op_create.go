// Package sync_lists contains auto-generated files. DO NOT MODIFY
package sync_lists

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateSyncListInput defines the input fields for creating a new list resource
type CreateSyncListInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
	UniqueName    *string `form:"UniqueName,omitempty"`
}

// CreateSyncListResponse defines the response fields for the created list
type CreateSyncListResponse struct {
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

// Create creates a new list
// See https://www.twilio.com/docs/sync/api/list-resource#create-a-list-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncListInput) (*CreateSyncListResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new list
// See https://www.twilio.com/docs/sync/api/list-resource#create-a-list-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncListInput) (*CreateSyncListResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateSyncListInput{}
	}

	response := &CreateSyncListResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
