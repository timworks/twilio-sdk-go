// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSyncMapPermissionsInput defines input fields for updating a map item permission resource
type UpdateSyncMapPermissionsInput struct {
	Manage bool `form:"Manage"`
	Read   bool `form:"Read"`
	Write  bool `form:"Write"`
}

// UpdateSyncMapPermissionsResponse defines the response fields for the updated map item permission
type UpdateSyncMapPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	Manage     bool   `json:"manage"`
	MapSid     string `json:"map_sid"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// Update modifies an map item permission resource
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#update-a-sync-map-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSyncMapPermissionsInput) (*UpdateSyncMapPermissionsResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies an map item permission resource
// See https://www.twilio.com/docs/sync/api/sync-map-permission-resource#update-a-sync-map-permission-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSyncMapPermissionsInput) (*UpdateSyncMapPermissionsResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Maps/{syncMapSid}/Permissions/{identity}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
			"identity":   c.identity,
		},
	}

	if input == nil {
		input = &UpdateSyncMapPermissionsInput{}
	}

	response := &UpdateSyncMapPermissionsResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
