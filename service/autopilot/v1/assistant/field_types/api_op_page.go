// Package field_types contains auto-generated files. DO NOT MODIFY
package field_types

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FieldTypesPageOptions defines the query options for the api operation
type FieldTypesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFieldTypeResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

// FieldTypesPageResponse defines the response fields for the field types page
type FieldTypesPageResponse struct {
	FieldTypes []PageFieldTypeResponse `json:"field_types"`
	Meta       PageMetaResponse        `json:"meta"`
}

// Page retrieves a page of field types
// See https://www.twilio.com/docs/autopilot/api/field-type#read-multiple-fieldtype-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FieldTypesPageOptions) (*FieldTypesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of field types
// See https://www.twilio.com/docs/autopilot/api/field-type#read-multiple-fieldtype-resources for more details
func (c Client) PageWithContext(context context.Context, options *FieldTypesPageOptions) (*FieldTypesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FieldTypesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FieldTypesPaginator defines the fields for makings paginated api calls
// FieldTypes is an array of fieldtypes that have been returned from all of the page calls
type FieldTypesPaginator struct {
	options    *FieldTypesPageOptions
	Page       *FieldTypesPage
	FieldTypes []PageFieldTypeResponse
}

// NewFieldTypesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFieldTypesPaginator() *FieldTypesPaginator {
	return c.NewFieldTypesPaginatorWithOptions(nil)
}

// NewFieldTypesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFieldTypesPaginatorWithOptions(options *FieldTypesPageOptions) *FieldTypesPaginator {
	return &FieldTypesPaginator{
		options: options,
		Page: &FieldTypesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		FieldTypes: make([]PageFieldTypeResponse, 0),
	}
}

// FieldTypesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFieldTypeResponse or error that is returned from the api call(s)
type FieldTypesPage struct {
	client *Client

	CurrentPage *FieldTypesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FieldTypesPaginator) CurrentPage() *FieldTypesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FieldTypesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FieldTypesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FieldTypesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FieldTypesPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().Meta.NextPageURL

		if nextPage == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPage)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.PageToken = utils.String(parsedURL.Query().Get("PageToken"))

		page, pageErr := strconv.Atoi(parsedURL.Query().Get("Page"))
		if pageErr != nil {
			p.Page.Error = pageErr
			return false
		}
		options.Page = utils.Int(page)

		pageSize, pageSizeErr := strconv.Atoi(parsedURL.Query().Get("PageSize"))
		if pageSizeErr != nil {
			p.Page.Error = pageSizeErr
			return false
		}
		options.PageSize = utils.Int(pageSize)
	}

	resp, err := p.Page.client.PageWithContext(context, options)
	p.Page.CurrentPage = resp
	p.Page.Error = err

	if p.Page.Error == nil {
		p.FieldTypes = append(p.FieldTypes, resp.FieldTypes...)
	}

	return p.Page.Error == nil
}
