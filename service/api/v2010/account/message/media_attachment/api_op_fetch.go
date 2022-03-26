// Package media_attachment contains auto-generated files. DO NOT MODIFY
package media_attachment

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchMediaResponse defines the response fields for the retrieved media attachment
type FetchMediaResponse struct {
	AccountSid  string             `json:"account_sid"`
	ContentType string             `json:"content_type"`
	DateCreated utils.RFC2822Time  `json:"date_created"`
	DateUpdated *utils.RFC2822Time `json:"date_updated,omitempty"`
	ParentSid   string             `json:"parent_sid"`
	Sid         string             `json:"sid"`
}

// Fetch retrieves a media attachment resource
// See https://www.twilio.com/docs/sms/api/media-resource#fetch-a-media-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchMediaResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a media attachment resource
// See https://www.twilio.com/docs/sms/api/media-resource#fetch-a-media-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchMediaResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Messages/{messageSid}/Media/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"messageSid": c.messageSid,
			"sid":        c.sid,
		},
	}

	response := &FetchMediaResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
