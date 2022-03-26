// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a address participant from the conference
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#delete-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a address participant from the conference
// See https://www.twilio.com/docs/voice/api/conference-participant-resource#delete-a-participant-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/Conferences/{conferenceSid}/Participants/{sid}.json",
		PathParams: map[string]string{
			"accountSid":    c.accountSid,
			"conferenceSid": c.conferenceSid,
			"sid":           c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
