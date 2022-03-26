// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
)

// Delete removes a credential list from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#delete-a-sip-credentiallist-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a credential list from the account
// See https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource#delete-a-sip-credentiallist-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Accounts/{accountSid}/SIP/CredentialLists/{sid}.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
			"sid":        c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
