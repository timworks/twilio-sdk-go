// Package members contains auto-generated files. DO NOT MODIFY
package members

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateChannelMemberInput defines the input fields for creating a new member resource
type CreateChannelMemberInput struct {
	Attributes               *string    `form:"Attributes,omitempty"`
	DateCreated              *time.Time `form:"DateCreated,omitempty"`
	DateUpdated              *time.Time `form:"DateUpdated,omitempty"`
	Identity                 string     `validate:"required" form:"Identity"`
	LastConsumedMessageIndex *int       `form:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `form:"LastConsumptionTimestamp,omitempty"`
	RoleSid                  *string    `form:"RoleSid,omitempty"`
}

// CreateChannelMemberResponse defines the response fields for the created member
type CreateChannelMemberResponse struct {
	AccountSid               string     `json:"account_sid"`
	Attributes               *string    `json:"attributes,omitempty"`
	ChannelSid               string     `json:"channel_sid"`
	DateCreated              time.Time  `json:"date_created"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	Identity                 string     `json:"identity"`
	LastConsumedMessageIndex *int       `json:"last_consumed_message_index,omitempty"`
	LastConsumedTimestamp    *time.Time `json:"last_consumption_timestamp,omitempty"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	ServiceSid               string     `json:"service_sid"`
	Sid                      string     `json:"sid"`
	URL                      string     `json:"url"`
}

// Create creates a new member
// See https://www.twilio.com/docs/chat/rest/member-resource#create-a-member-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateChannelMemberInput) (*CreateChannelMemberResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new member
// See https://www.twilio.com/docs/chat/rest/member-resource#create-a-member-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateChannelMemberInput) (*CreateChannelMemberResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Channels/{channelSid}/Members",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
	}

	if input == nil {
		input = &CreateChannelMemberInput{}
	}

	response := &CreateChannelMemberResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
