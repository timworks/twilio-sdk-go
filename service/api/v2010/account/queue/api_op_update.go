// Package queue contains auto-generated files. DO NOT MODIFY
package queue

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UpdateQueueInput defines input fields for updating a queue
type UpdateQueueInput struct {
	FriendlyName *string `form:"FriendlyName,omitempty"`
	MaxSize      *int    `form:"MaxSize,omitempty"`
}

// UpdateQueueResponse defines the response fields for the updated queue
type UpdateQueueResponse struct {
	AccountSid      string             `json:"account_sid"`
	AverageWaitTime int                `json:"average_wait_time"`
	CurrentSize     int                `json:"current_size"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	MaxSize         int                `json:"max_size"`
	Sid             string             `json:"sid"`
}

// Update modifies a queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#update-a-queue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateQueueInput) (*UpdateQueueResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#update-a-queue-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateQueueInput) (*UpdateQueueResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Accounts/{accountSid}/Queues/{sid}.json",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateQueueInput{}
	}

	response := &UpdateQueueResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
