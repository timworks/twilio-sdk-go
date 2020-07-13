// This is an autogenerated file. DO NOT MODIFY
package member

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type GetChannelMemberOutput struct {
	Sid                      string     `json:"sid"`
	AccountSid               string     `json:"account_sid"`
	ServiceSid               string     `json:"service_sid"`
	ChannelSid               string     `json:"channel_sid"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	Identity                 string     `json:"identity"`
	LastConsumedMessageIndex *int       `json:"last_consumed_message_index,omitempty"`
	LastConsumedTimestamp    *time.Time `json:"last_consumption_timestamp,omitempty"`
	Attributes               *string    `json:"attributes,omitempty"`
	DateCreated              time.Time  `json:"date_created"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	URL                      string     `json:"url"`
}

func (c Client) Get() (*GetChannelMemberOutput, error) {
	return c.GetWithContext(context.Background())
}

func (c Client) GetWithContext(context context.Context) (*GetChannelMemberOutput, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
			"sid":        c.sid,
		},
	}

	output := &GetChannelMemberOutput{}
	if err := c.client.Send(context, op, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}
