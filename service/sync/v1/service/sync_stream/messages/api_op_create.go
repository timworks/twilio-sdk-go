// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateSyncStreamMessageInput defines the input fields for creating a new stream message resource
type CreateSyncStreamMessageInput struct {
	Data string `validate:"required" form:"Data"`
}

// CreateSyncStreamMessageResponse defines the response fields for the created stream message
type CreateSyncStreamMessageResponse struct {
	Data map[string]interface{} `json:"data"`
	Sid  string                 `json:"sid"`
}

// Create creates a new stream message
// See https://www.twilio.com/docs/sync/api/stream-message-resource#create-a-stream-message-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateSyncStreamMessageInput) (*CreateSyncStreamMessageResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new stream message
// See https://www.twilio.com/docs/sync/api/stream-message-resource#create-a-stream-message-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateSyncStreamMessageInput) (*CreateSyncStreamMessageResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Streams/{syncStreamSid}/Messages",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid":    c.serviceSid,
			"syncStreamSid": c.syncStreamSid,
		},
	}

	if input == nil {
		input = &CreateSyncStreamMessageInput{}
	}

	response := &CreateSyncStreamMessageResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
