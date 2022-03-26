// Package user contains auto-generated files. DO NOT MODIFY
package user

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchUserResponse defines the response fields for the retrieved user
type FetchUserResponse struct {
	AccountSid          string     `json:"account_sid"`
	Attributes          *string    `json:"attributes,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	FriendlyName        *string    `json:"friendly_name,omitempty"`
	Identity            string     `json:"identity"`
	IsNotifiable        *bool      `json:"is_notifiable,omitempty"`
	IsOnline            *bool      `json:"is_online,omitempty"`
	JoinedChannelsCount *int       `json:"joined_channels_count,omitempty"`
	RoleSid             string     `json:"role_sid"`
	ServiceSid          string     `json:"service_sid"`
	Sid                 string     `json:"sid"`
	URL                 string     `json:"url"`
}

// Fetch retrieves a user resource
// See https://www.twilio.com/docs/chat/rest/user-resource#fetch-a-user-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchUserResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a user resource
// See https://www.twilio.com/docs/chat/rest/user-resource#fetch-a-user-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchUserResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchUserResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
