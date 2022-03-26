// Package invites contains auto-generated files. DO NOT MODIFY
package invites

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateChannelInviteInput defines the input fields for creating a new invite resource
type CreateChannelInviteInput struct {
	Identity string  `validate:"required" form:"Identity"`
	RoleSid  *string `form:"RoleSid,omitempty"`
}

// CreateChannelInviteResponse defines the response fields for the created invite
type CreateChannelInviteResponse struct {
	AccountSid  string     `json:"account_sid"`
	ChannelSid  string     `json:"channel_sid"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    string     `json:"identity"`
	RoleSid     *string    `json:"role_sid,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Create creates a new invite
// See https://www.twilio.com/docs/chat/rest/invite-resource#create-an-invite-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelInviteInput) (*CreateChannelInviteResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new invite
// See https://www.twilio.com/docs/chat/rest/invite-resource#create-an-invite-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChannelInviteInput) (*CreateChannelInviteResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Invites",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	if input == nil {
		input = &CreateChannelInviteInput{}
	}

	response := &CreateChannelInviteResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
