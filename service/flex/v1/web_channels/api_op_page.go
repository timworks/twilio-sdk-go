// Package web_channels contains auto-generated files. DO NOT MODIFY
package web_channels

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// WebChannelsPageOptions defines the query options for the api operation
type WebChannelsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
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

type PageWebChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
}

// WebChannelsPageResponse defines the response fields for the web channels page
type WebChannelsPageResponse struct {
	Meta        PageMetaResponse         `json:"meta"`
	WebChannels []PageWebChannelResponse `json:"flex_chat_channels"`
}

// Page retrieves a page of web channels
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *WebChannelsPageOptions) (*WebChannelsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of web channels
func (c Client) PageWithContext(context context.Context, options *WebChannelsPageOptions) (*WebChannelsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/WebChannels",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &WebChannelsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// WebChannelsPaginator defines the fields for makings paginated api calls
// WebChannels is an array of webchannels that have been returned from all of the page calls
type WebChannelsPaginator struct {
	options     *WebChannelsPageOptions
	Page        *WebChannelsPage
	WebChannels []PageWebChannelResponse
}

// NewWebChannelsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewWebChannelsPaginator() *WebChannelsPaginator {
	return c.NewWebChannelsPaginatorWithOptions(nil)
}

// NewWebChannelsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewWebChannelsPaginatorWithOptions(options *WebChannelsPageOptions) *WebChannelsPaginator {
	return &WebChannelsPaginator{
		options: options,
		Page: &WebChannelsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		WebChannels: make([]PageWebChannelResponse, 0),
	}
}

// WebChannelsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageWebChannelResponse or error that is returned from the api call(s)
type WebChannelsPage struct {
	client *Client

	CurrentPage *WebChannelsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *WebChannelsPaginator) CurrentPage() *WebChannelsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *WebChannelsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *WebChannelsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *WebChannelsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &WebChannelsPageOptions{}
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
		p.WebChannels = append(p.WebChannels, resp.WebChannels...)
	}

	return p.Page.Error == nil
}
