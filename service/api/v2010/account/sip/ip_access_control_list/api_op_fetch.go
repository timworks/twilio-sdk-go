// Package ip_access_control_list contains auto-generated files. DO NOT MODIFY
package ip_access_control_list

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchIpAccessControlListResponse defines the response fields for retrieving a credential list
type FetchIpAccessControlListResponse struct {
	AccountSid   string             `json:"account_sid"`
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// Fetch retrieves a IP control access list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#fetch-a-sip-ipaccesscontrollist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchIpAccessControlListResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a IP control access list resource
// See https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource#fetch-a-sip-ipaccesscontrollist-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchIpAccessControlListResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/SIP/IpAccessControlLists/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchIpAccessControlListResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
