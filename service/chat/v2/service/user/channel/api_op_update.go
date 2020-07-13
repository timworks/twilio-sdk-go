// This is an autogenerated file. DO NOT MODIFY
package channel

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateUserChannelInput struct {
	NotificationLevel        *string    `form:"NotificationLevel,omitempty"`
	LastConsumedMessageIndex *int       `form:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp *time.Time `form:"LastConsumptionTimestamp,omitempty"`
}

type UpdateUserChannelOutput struct {
	AccountSid               string  `json:"account_sid"`
	ServiceSid               string  `json:"service_sid"`
	ChannelSid               string  `json:"channel_sid"`
	UserSid                  string  `json:"user_sid"`
	MemberSid                string  `json:"member_sid"`
	Status                   string  `json:"status"`
	LastConsumedMessageIndex *int    `json:"last_consumed_message_index,omitempty"`
	UnreadMessagesCount      *int    `json:"unread_messages_count,omitempty"`
	NotificationLevel        *string `json:"notification_level,omitempty"`
	URL                      string  `json:"url"`
}

func (c Client) Update(input *UpdateUserChannelInput) (*UpdateUserChannelOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateUserChannelInput) (*UpdateUserChannelOutput, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Users/{userSid}/Channels/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
			"sid":        c.sid,
		},
	}

	output := &UpdateUserChannelOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
