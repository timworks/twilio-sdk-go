// Package phone_number contains auto-generated files. DO NOT MODIFY
package phone_number

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// FetchPhoneNumberResponse defines the response fields for the retrieved phone number
type FetchPhoneNumberResponse struct {
	AccountSid   string     `json:"account_sid"`
	Capabilities []string   `json:"capabilities"`
	CountryCode  string     `json:"country_code"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	PhoneNumber  string     `json:"phone_number"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
}

// Fetch retrieves a phone number resource
// See https://www.twilio.com/docs/sms/services/api/phonenumber-resource#fetch-a-phonenumber-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch() (*FetchPhoneNumberResponse, error) {
	return c.FetchWithContext(context.Background())
}

// FetchWithContext retrieves a phone number resource
// See https://www.twilio.com/docs/sms/services/api/phonenumber-resource#fetch-a-phonenumber-resource for more details
func (c Client) FetchWithContext(context context.Context) (*FetchPhoneNumberResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/PhoneNumbers/{sid}",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"sid":        c.sid,
		},
	}

	response := &FetchPhoneNumberResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
