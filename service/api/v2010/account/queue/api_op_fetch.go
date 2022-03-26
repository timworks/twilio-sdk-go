// Package queue contains auto-generated files. DO NOT MODIFY
package queue

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchQueueResponse defines the response fields for retrieving a queue
type FetchQueueResponse struct {
	AccountSid      string             `json:"account_sid"`
	AverageWaitTime int                `json:"average_wait_time"`
	CurrentSize     int                `json:"current_size"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	MaxSize         int                `json:"max_size"`
	Sid             string             `json:"sid"`
}

// Fetch retrieves the queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#fetch-a-queue-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchQueueResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves the queue resource
// See https://www.twilio.com/docs/voice/api/queue-resource#fetch-a-queue-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchQueueResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Queues/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	response := &FetchQueueResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
