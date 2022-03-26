// Package item contains auto-generated files. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSyncMapItemInput defines input fields for updating a map item resource
type UpdateSyncMapItemInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Data          *string `form:"Data,omitempty"`
	ItemTtl       *int    `form:"ItemTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
}

// UpdateSyncMapItemResponse defines the response fields for the updated map item
type UpdateSyncMapItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Key         string                 `json:"key"`
	MapSid      string                 `json:"map_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

// Update modifies an map item resource
// See https://www.twilio.com/docs/sync/api/map-item-resource#update-a-mapitem-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncMapItemInput) (*UpdateSyncMapItemResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an map item resource
// See https://www.twilio.com/docs/sync/api/map-item-resource#update-a-mapitem-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncMapItemInput) (*UpdateSyncMapItemResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Maps/{syncMapSid}/Items/{key}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
			"key":        c.key,
		},
	}

	if input == nil {
		input = &UpdateSyncMapItemInput{}
	}

	response := &UpdateSyncMapItemResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
