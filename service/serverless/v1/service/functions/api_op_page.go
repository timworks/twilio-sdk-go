// Package functions contains auto-generated files. DO NOT MODIFY
package functions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// FunctionsPageOptions defines the query options for the api operation
type FunctionsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageFunctionResponse struct {
	AccountSid   string     `json:"account_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName string     `json:"friendly_name"`
	ServiceSid   string     `json:"service_sid"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
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

// FunctionsPageResponse defines the response fields for the functions page
type FunctionsPageResponse struct {
	Functions []PageFunctionResponse `json:"functions"`
	Meta      PageMetaResponse       `json:"meta"`
}

// Page retrieves a page of functions
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#read-multiple-function-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *FunctionsPageOptions) (*FunctionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of functions
// See https://www.twilio.com/docs/runtime/functions-assets-api/api/function#read-multiple-function-resources for more details
func (c Client) PageWithContext(context context.Context, options *FunctionsPageOptions) (*FunctionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Functions",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FunctionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// FunctionsPaginator defines the fields for makings paginated api calls
// Functions is an array of functions that have been returned from all of the page calls
type FunctionsPaginator struct {
	options   *FunctionsPageOptions
	Page      *FunctionsPage
	Functions []PageFunctionResponse
}

// NewFunctionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewFunctionsPaginator() *FunctionsPaginator {
	return c.NewFunctionsPaginatorWithOptions(nil)
}

// NewFunctionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewFunctionsPaginatorWithOptions(options *FunctionsPageOptions) *FunctionsPaginator {
	return &FunctionsPaginator{
		options: options,
		Page: &FunctionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Functions: make([]PageFunctionResponse, 0),
	}
}

// FunctionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageFunctionResponse or error that is returned from the api call(s)
type FunctionsPage struct {
	client *Client

	CurrentPage *FunctionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *FunctionsPaginator) CurrentPage() *FunctionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *FunctionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *FunctionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *FunctionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &FunctionsPageOptions{}
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
		p.Functions = append(p.Functions, resp.Functions...)
	}

	return p.Page.Error == nil
}
