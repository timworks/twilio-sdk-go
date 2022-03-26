// Package asset contains auto-generated files. DO NOT MODIFY
package asset

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a asset resource from the account
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#delete-an-asset-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a asset resource from the account
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/asset#delete-an-asset-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Assets/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
