// Package permission contains auto-generated files. DO NOT MODIFY
package permission

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchSyncListPermissionsResponse defines the response fields for the retrieved list permission
type FetchSyncListPermissionsResponse struct {
	AccountSid string `json:"account_sid"`
	Identity   string `json:"identity"`
	ListSid    string `json:"list_sid"`
	Manage     bool   `json:"manage"`
	Read       bool   `json:"read"`
	ServiceSid string `json:"service_sid"`
	URL        string `json:"url"`
	Write      bool   `json:"write"`
}

// Fetch retrieves an list permission resource
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#fetch-a-sync-list-permission-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchSyncListPermissionsResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves an list permission resource
// See https://www.twilio.com/docs/sync/api/sync-list-permission-resource#fetch-a-sync-list-permission-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchSyncListPermissionsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Lists/{syncListSid}/Permissions/{identity}",
		PathParams: map[string]string{
			"serviceSid":  c.serviceSid,
			"syncListSid": c.syncListSid,
			"identity":    c.identity,
		},
	}

	response := &FetchSyncListPermissionsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
