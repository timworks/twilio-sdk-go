// Package alpha_senders contains auto-generated files. DO NOT MODIFY
package alpha_senders

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// AlphaSendersPageOptions defines the query options for the api operation
type AlphaSendersPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageAlphaSenderResponse struct {
	AccountSid   string     `json:"account_sid"`
	AlphaSender  string     `json:"alpha_sender"`
	Capabilities []string   `json:"capabilities"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
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

// AlphaSendersPageResponse defines the response fields for the alpha senders page
type AlphaSendersPageResponse struct {
	AlphaSenders []PageAlphaSenderResponse `json:"alpha_senders"`
	Meta         PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of alpha senders
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#read-multiple-alphasender-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AlphaSendersPageOptions) (*AlphaSendersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of alpha senders
// See https://www.twilio.com/docs/sms/services/api/alphasender-resource#read-multiple-alphasender-resources for more details
func (c Client) PageWithContext(context context.Context, options *AlphaSendersPageOptions) (*AlphaSendersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/AlphaSenders",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AlphaSendersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AlphaSendersPaginator defines the fields for makings paginated api calls
// AlphaSenders is an array of alphasenders that have been returned from all of the page calls
type AlphaSendersPaginator struct {
	options      *AlphaSendersPageOptions
	Page         *AlphaSendersPage
	AlphaSenders []PageAlphaSenderResponse
}

// NewAlphaSendersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAlphaSendersPaginator() *AlphaSendersPaginator {
	return c.NewAlphaSendersPaginatorWithOptions(nil)
}

// NewAlphaSendersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAlphaSendersPaginatorWithOptions(options *AlphaSendersPageOptions) *AlphaSendersPaginator {
	return &AlphaSendersPaginator{
		options: options,
		Page: &AlphaSendersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		AlphaSenders: make([]PageAlphaSenderResponse, 0),
	}
}

// AlphaSendersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAlphaSenderResponse or error that is returned from the api call(s)
type AlphaSendersPage struct {
	client *Client

	CurrentPage *AlphaSendersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AlphaSendersPaginator) CurrentPage() *AlphaSendersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AlphaSendersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AlphaSendersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AlphaSendersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AlphaSendersPageOptions{}
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
		p.AlphaSenders = append(p.AlphaSenders, resp.AlphaSenders...)
	}

	return p.Page.Error == nil
}
