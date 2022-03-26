// Package faxes contains auto-generated files. DO NOT MODIFY
package faxes

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
)

// CreateFaxInput defines the input fields for creating a new fax resource
type CreateFaxInput struct {
	From            *string `form:"From,omitempty"`
	MediaURL        string  `validate:"required" form:"MediaUrl"`
	Quality         *string `form:"Quality,omitempty"`
	SipAuthPassword *string `form:"SipAuthPassword,omitempty"`
	SipAuthUsername *string `form:"SipAuthUsername,omitempty"`
	StatusCallback  *string `form:"StatusCallback,omitempty"`
	StoreMedia      *bool   `form:"StoreMedia,omitempty"`
	To              string  `validate:"required" form:"To"`
	Ttl             *int    `form:"Ttl,omitempty"`
}

// CreateFaxResponse defines the response fields for the created fax
type CreateFaxResponse struct {
	APIVersion  string     `json:"api_version"`
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Direction   string     `json:"direction"`
	Duration    *int       `json:"duration,omitempty"`
	From        string     `json:"from"`
	MediaSid    *string    `json:"media_sid,omitempty"`
	MediaURL    *string    `json:"media_url,omitempty"`
	NumPages    *int       `json:"num_pages,omitempty"`
	Price       *string    `json:"price,omitempty"`
	PriceUnit   *string    `json:"price_unit,omitempty"`
	Quality     string     `json:"quality"`
	Sid         string     `json:"sid"`
	Status      string     `json:"status"`
	To          string     `json:"to"`
	URL         string     `json:"url"`
}

// Create creates a new fax
// See https://www.twilio.com/docs/fax/api/fax-resource#create-a-fax-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Create(input *CreateFaxInput) (*CreateFaxResponse, error) {
	return c.CreateWithContext(context.Background(), input)
}

// CreateWithContext creates a new fax
// See https://www.twilio.com/docs/fax/api/fax-resource#create-a-fax-resource for more details
func (c Client) CreateWithContext(context context.Context, input *CreateFaxInput) (*CreateFaxResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Faxes",
		ContentType: client.URLEncoded,
	}

	if input == nil {
		input = &CreateFaxInput{}
	}

	response := &CreateFaxResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}
