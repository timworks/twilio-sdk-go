// Package revision contains auto-generated files. DO NOT MODIFY
package revision

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchRevisionResponse defines the response fields for the retrieved flow revision
type FetchRevisionResponse struct {
	AccountSid    string                 `json:"account_sid"`
	CommitMessage *string                `json:"commit_message,omitempty"`
	DateCreated   time.Time              `json:"date_created"`
	DateUpdated   *time.Time             `json:"date_updated,omitempty"`
	Definition    map[string]interface{} `json:"definition"`
	Errors        *[]interface{}         `json:"errors,omitempty"`
	FriendlyName  string                 `json:"friendly_name"`
	Revision      int                    `json:"revision"`
	Sid           string                 `json:"sid"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Valid         bool                   `json:"valid"`
	Warnings      *[]interface{}         `json:"warnings,omitempty"`
	WebhookURL    string                 `json:"webhook_url"`
}

// Fetch retrieves a flow revision resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision#fetch-a-flowrevision-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchRevisionResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a flow revision resource
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision#fetch-a-flowrevision-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchRevisionResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Revisions/{revisionNumber}",
		PathParams: map[string]string{
			"flowSid":        c.flowSid,
			"revisionNumber": strconv.Itoa(c.revisionNumber),
		},
	}

	response := &FetchRevisionResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
