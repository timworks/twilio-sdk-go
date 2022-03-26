// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// ChannelsPageOptions defines the query options for the api operation
type ChannelsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageChannelResponse struct {
	AccountSid  string     `json:"account_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	Sid         string     `json:"sid"`
	TaskSid     *string    `json:"task_sid,omitempty"`
	URL         string     `json:"url"`
	UserSid     string     `json:"user_sid"`
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

// ChannelsPageResponse defines the response fields for the channels page
type ChannelsPageResponse struct {
	Channels []PageChannelResponse `json:"flex_chat_channels"`
	Meta     PageMetaResponse      `json:"meta"`
}

// Page retrieves a page of channels
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ChannelsPageOptions) (*ChannelsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of channels
func (c Client) PageWithContext(context context.Context, options *ChannelsPageOptions) (*ChannelsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Channels",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChannelsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChannelsPaginator defines the fields for makings paginated api calls
// Channels is an array of channels that have been returned from all of the page calls
type ChannelsPaginator struct {
	options  *ChannelsPageOptions
	Page     *ChannelsPage
	Channels []PageChannelResponse
}

// NewChannelsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChannelsPaginator() *ChannelsPaginator {
	return c.NewChannelsPaginatorWithOptions(nil)
}

// NewChannelsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChannelsPaginatorWithOptions(options *ChannelsPageOptions) *ChannelsPaginator {
	return &ChannelsPaginator{
		options: options,
		Page: &ChannelsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Channels: make([]PageChannelResponse, 0),
	}
}

// ChannelsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChannelResponse or error that is returned from the api call(s)
type ChannelsPage struct {
	client *Client

	CurrentPage *ChannelsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChannelsPaginator) CurrentPage() *ChannelsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChannelsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChannelsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChannelsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChannelsPageOptions{}
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
		p.Channels = append(p.Channels, resp.Channels...)
	}

	return p.Page.Error == nil
}
