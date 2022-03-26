// Package web_channel contains auto-generated files. DO NOT MODIFY
package web_channel

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateWebChannelInput defines input fields for updating a web channel resource
type UpdateWebChannelInput struct {
	ChatStatus        *string `form:"ChatStatus,omitempty"`
	PreEngagementData *string `form:"PreEngagementData,omitempty"`
}

// UpdateWebChannelResponse defines the response fields for the updated web channel
type UpdateWebChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// Update modifies a web channel resource
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateWebChannelInput) (*UpdateWebChannelResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a web channel resource
func (c Client) UpdateWithContext(context context.Context, input *UpdateWebChannelInput) (*UpdateWebChannelResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/WebChannels/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if input == nil {
		input = &UpdateWebChannelInput{}
	}

	response := &UpdateWebChannelResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
