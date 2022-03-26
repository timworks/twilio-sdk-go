// Package channel contains auto-generated files. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchUserChannelResponse defines the response fields for the retrieved user channel
type FetchUserChannelResponse struct {
	AccountSid               string  `json:"account_sid"`
	ChannelSid               string  `json:"channel_sid"`
	LastConsumedMessageIndex *int    `json:"last_consumed_message_index,omitempty"`
	MemberSid                string  `json:"member_sid"`
	NotificationLevel        *string `json:"notification_level,omitempty"`
	ServiceSid               string  `json:"service_sid"`
	Status                   string  `json:"status"`
	URL                      string  `json:"url"`
	UnreadMessagesCount      *int    `json:"unread_messages_count,omitempty"`
	UserSid                  string  `json:"user_sid"`
}

// Fetch retrieves a user channel resource
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#fetch-a-userchannel-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchUserChannelResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a user channel resource
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#fetch-a-userchannel-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchUserChannelResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users/{userSid}/Channels/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
			"sid":        c.sid,
		},
	}

	response := &FetchUserChannelResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
