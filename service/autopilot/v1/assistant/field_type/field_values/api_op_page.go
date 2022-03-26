// Package field_values contains auto-generated files. DO NOT MODIFY
package field_values

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FieldValuesPageOptions defines the query options for the api operation
type FieldValuesPageOptions struct {
	Language  *string
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFieldValueResponse struct {
	AccountSid   string     `json:"account_sid"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FieldTypeSid string     `json:"field_type_sid"`
	Language     string     `json:"language"`
	Sid          string     `json:"sid"`
	SynonymOf    *string    `json:"synonym_of,omitempty"`
	URL          string     `json:"url"`
	Value        string     `json:"value"`
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

// FieldValuesPageResponse defines the response fields for the field value page
type FieldValuesPageResponse struct {
	FieldValues []PageFieldValueResponse `json:"field_values"`
	Meta        PageMetaResponse         `json:"meta"`
}

// Page retrieves a page of field values
// See https://www.twilio.com/docs/autopilot/api/field-value#read-multiple-fieldvalue-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FieldValuesPageOptions) (*FieldValuesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of field values
// See https://www.twilio.com/docs/autopilot/api/field-value#read-multiple-fieldvalue-resources for more details
func (c Client) PageWithContext(context context.Context, options *FieldValuesPageOptions) (*FieldValuesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
			"fieldTypeSid": c.fieldTypeSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FieldValuesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FieldValuesPaginator defines the fields for makings paginated api calls
// FieldValues is an array of fieldvalues that have been returned from all of the page calls
type FieldValuesPaginator struct {
	options     *FieldValuesPageOptions
	Page        *FieldValuesPage
	FieldValues []PageFieldValueResponse
}

// NewFieldValuesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFieldValuesPaginator() *FieldValuesPaginator {
	return c.NewFieldValuesPaginatorWithOptions(nil)
}

// NewFieldValuesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFieldValuesPaginatorWithOptions(options *FieldValuesPageOptions) *FieldValuesPaginator {
	return &FieldValuesPaginator{
		options: options,
		Page: &FieldValuesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		FieldValues: make([]PageFieldValueResponse, 0),
	}
}

// FieldValuesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFieldValueResponse or error that is returned from the api call(s)
type FieldValuesPage struct {
	client *Client

	CurrentPage *FieldValuesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FieldValuesPaginator) CurrentPage() *FieldValuesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FieldValuesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FieldValuesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FieldValuesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FieldValuesPageOptions{}
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
		p.FieldValues = append(p.FieldValues, resp.FieldValues...)
	}

	return p.Page.Error == nil
}
