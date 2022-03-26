// Package composition contains auto-generated files. DO NOT MODIFY
package composition

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a composition resource from the account
// See https://www.twilio.com/docs/video/api/compositions-resource#delete-composition-http-delete for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a composition resource from the account
// See https://www.twilio.com/docs/video/api/compositions-resource#delete-composition-http-delete for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Compositions/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
