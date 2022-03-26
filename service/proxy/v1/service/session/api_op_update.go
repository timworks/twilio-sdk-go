// Package session contains auto-generated files. DO NOT MODIFY
package session

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// UpdateSessionInput defines input fields for updating a session resource
type UpdateSessionInput struct {
	DateExpiry *time.Time `form:"DateExpiry,omitempty"`
	Status     *string    `form:"Status,omitempty"`
	Ttl        *int       `form:"Ttl,omitempty"`
}

// UpdateSessionResponse defines the response fields for the updated session
type UpdateSessionResponse struct {
	AccountSid          string     `json:"account_sid"`
	ClosedReason        *string    `json:"closed_reason,omitempty"`
	DateCreated         time.Time  `json:"date_created"`
	DateEnded           *time.Time `json:"date_ended,omitempty"`
	DateExpiry          *time.Time `json:"date_expiry,omitempty"`
	DateLastInteraction *time.Time `json:"date_last_interaction,omitempty"`
	DateStarted         *time.Time `json:"date_started,omitempty"`
	DateUpdated         *time.Time `json:"date_updated,omitempty"`
	Mode                *string    `json:"mode,omitempty"`
	ServiceSid          string     `json:"service_sid"`
	Sid                 string     `json:"sid"`
	Status              *string    `json:"status,omitempty"`
	Ttl                 *int       `json:"ttl,omitempty"`
	URL                 string     `json:"url"`
	UniqueName          string     `json:"unique_name"`
}

// Update modifies a session resource
// See https://www.twilio.com/docs/proxy/api/session#update-a-session-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Update(input *UpdateSessionInput) (*UpdateSessionResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

// UpdateWithContext modifies a session resource
// See https://www.twilio.com/docs/proxy/api/session#update-a-session-resource for more details
func (c Client) UpdateWithContext(context context.Context, input *UpdateSessionInput) (*UpdateSessionResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Sessions/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if input == nil {
		input = &UpdateSessionInput{}
	}

	response := &UpdateSessionResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
