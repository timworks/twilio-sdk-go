// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a participant resource from the account
// See https://www.twilio.com/docs/proxy/api/participant#delete-a-participant-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a participant resource from the account
// See https://www.twilio.com/docs/proxy/api/participant#delete-a-participant-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Sessions/{sessionSid}/Participants/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sessionSid": c.sessionSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
