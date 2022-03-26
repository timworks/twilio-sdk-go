// Package environments contains auto-generated files. DO NOT MODIFY
package environments

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// EnvironmentsPageOptions defines the query options for the api operation
type EnvironmentsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageEnvironmentResponse struct {
	AccountSid   string     `json:"account_sid"`
	BuildSid     *string    `json:"build_sid,omitempty"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	DomainName   string     `json:"domain_name"`
	DomainSuffix *string    `json:"domain_suffix,omitempty"`
	ServiceSid   string     `json:"service_sid"`
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

// EnvironmentsPageResponse defines the response fields for the environments page
type EnvironmentsPageResponse struct {
	Environments []PageEnvironmentResponse `json:"environments"`
	Meta         PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of environments
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#read-multiple-environment-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *EnvironmentsPageOptions) (*EnvironmentsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of environments
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/environment#read-multiple-environment-resources for more details
func (c Client) PageWithContext(context context.Context, options *EnvironmentsPageOptions) (*EnvironmentsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Environments",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &EnvironmentsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// EnvironmentsPaginator defines the fields for makings paginated api calls
// Environments is an array of environments that have been returned from all of the page calls
type EnvironmentsPaginator struct {
	options      *EnvironmentsPageOptions
	Page         *EnvironmentsPage
	Environments []PageEnvironmentResponse
}

// NewEnvironmentsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewEnvironmentsPaginator() *EnvironmentsPaginator {
	return c.NewEnvironmentsPaginatorWithOptions(nil)
}

// NewEnvironmentsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewEnvironmentsPaginatorWithOptions(options *EnvironmentsPageOptions) *EnvironmentsPaginator {
	return &EnvironmentsPaginator{
		options: options,
		Page: &EnvironmentsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Environments: make([]PageEnvironmentResponse, 0),
	}
}

// EnvironmentsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageEnvironmentResponse or error that is returned from the api call(s)
type EnvironmentsPage struct {
	client *Client

	CurrentPage *EnvironmentsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *EnvironmentsPaginator) CurrentPage() *EnvironmentsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *EnvironmentsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *EnvironmentsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *EnvironmentsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &EnvironmentsPageOptions{}
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
		p.Environments = append(p.Environments, resp.Environments...)
	}

	return p.Page.Error == nil
}
