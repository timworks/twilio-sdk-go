// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateChannelInput defines the input fields for creating a new channel resource
type CreateChannelInput struct {
	Attributes   *string    `form:"Attributes,omitempty"`
	CreatedBy    *string    `form:"CreatedBy,omitempty"`
	DateCreated  *time.Time `form:"DateCreated,omitempty"`
	DateUpdated  *time.Time `form:"DateUpdated,omitempty"`
	FriendlyName *string    `form:"FriendlyName,omitempty"`
	Type         *string    `form:"Type,omitempty"`
	UniqueName   *string    `form:"UniqueName,omitempty"`
}

// CreateChannelResponse defines the response fields for the created channel
type CreateChannelResponse struct {
	AccountSid    string     `json:"account_sid"`
	Attributes    *string    `json:"attributes,omitempty"`
	CreatedBy     string     `json:"created_by"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	FriendlyName  *string    `json:"friendly_name,omitempty"`
	MembersCount  int        `json:"members_count"`
	MessagesCount int        `json:"messages_count"`
	ServiceSid    string     `json:"service_sid"`
	Sid           string     `json:"sid"`
	Type          string     `json:"type"`
	URL           string     `json:"url"`
	UniqueName    *string    `json:"unique_name,omitempty"`
}

// Create creates a new channel
// See https://www.twilio.com/docs/chat/rest/channel-resource#create-a-channel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelInput) (*CreateChannelResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new channel
// See https://www.twilio.com/docs/chat/rest/channel-resource#create-a-channel-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChannelInput) (*CreateChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	if input == nil {
		input = &CreateChannelInput{}
	}

	response := &CreateChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
