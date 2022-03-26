// Package media_attachment contains auto-generated files. DO NOT MODIFY
package media_attachment

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a media attachment resource from the account
// See https://www.twilio.com/docs/sms/api/media-resource#delete-a-media-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a media attachment resource from the account
// See https://www.twilio.com/docs/sms/api/media-resource#delete-a-media-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/Messages/{messageSid}/Media/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"messageSid": c.messageSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
