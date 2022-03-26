// Package binding contains auto-generated files. DO NOT MODIFY
package binding

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a user binding resource from the account
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#delete-a-user-binding-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a user binding resource from the account
// See https://www.twilio.com/docs/chat/rest/user-binding-resource#delete-a-user-binding-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Users/{userSid}/Bindings/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
