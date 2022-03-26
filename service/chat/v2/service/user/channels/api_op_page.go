// Package channels contains auto-generated files. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// UserChannelsPageOptions defines the query options for the api operation
type UserChannelsPageOptions struct {
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

type PageUserChannelResponse struct {
	AccountSid               string  `json:"account_sid"`
	ChannelSid               string  `json:"channel_sid"`
	LastConsumedMessageIndex *int    `json:"last_consumed_message_index,omitempty"`
	MemberSid                string  `json:"member_sid"`
	NotificationLevel        *string `json:"notification_level,omitempty"`
	ServiceSid               string  `json:"service_sid"`
	Status                   string  `json:"status"`
	URL                      string  `json:"url"`
	UnreadMessagesCount      *int    `json:"unread_messages_count,omitempty"`
	UserSid                  string  `json:"user_sid"`
}

// UserChannelsPageResponse defines the response fields for the user channels page
type UserChannelsPageResponse struct {
	Channels []PageUserChannelResponse `json:"channels"`
	Meta     PageMetaResponse          `json:"meta"`
}

// Page retrieves a page of user channels
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#read-multiple-userchannel-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *UserChannelsPageOptions) (*UserChannelsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of user channels
// See https://www.twilio.com/docs/chat/rest/user-channel-resource#read-multiple-userchannel-resources for more details
func (c Client) PageWithContext(context context.Context, options *UserChannelsPageOptions) (*UserChannelsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Users/{userSid}/Channels",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"userSid":    c.userSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &UserChannelsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// UserChannelsPaginator defines the fields for makings paginated api calls
// Channels is an array of channels that have been returned from all of the page calls
type UserChannelsPaginator struct {
	options  *UserChannelsPageOptions
	Page     *UserChannelsPage
	Channels []PageUserChannelResponse
}

// NewUserChannelsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewUserChannelsPaginator() *UserChannelsPaginator {
	return c.NewUserChannelsPaginatorWithOptions(nil)
}

// NewUserChannelsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewUserChannelsPaginatorWithOptions(options *UserChannelsPageOptions) *UserChannelsPaginator {
	return &UserChannelsPaginator{
		options: options,
		Page: &UserChannelsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Channels: make([]PageUserChannelResponse, 0),
	}
}

// UserChannelsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageUserChannelResponse or error that is returned from the api call(s)
type UserChannelsPage struct {
	client *Client

	CurrentPage *UserChannelsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *UserChannelsPaginator) CurrentPage() *UserChannelsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *UserChannelsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *UserChannelsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *UserChannelsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &UserChannelsPageOptions{}
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
