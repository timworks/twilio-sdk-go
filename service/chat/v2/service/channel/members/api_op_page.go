// Package members contains auto-generated files. DO NOT MODIFY
package members

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// ChannelMembersPageOptions defines the query options for the api operation
type ChannelMembersPageOptions struct {
	Identity  *[]string
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageChannelMemberResponse struct {
	AccountSid               string     `json:"account_sid"`
	Attributes               *string    `json:"attributes,omitempty"`
	ChannelSid               string     `json:"channel_sid"`
	DateCreated              time.Time  `json:"date_created"`
	DateUpdated              *time.Time `json:"date_updated,omitempty"`
	Identity                 string     `json:"identity"`
	LastConsumedMessageIndex *int       `json:"last_consumed_message_index,omitempty"`
	LastConsumedTimestamp    *time.Time `json:"last_consumption_timestamp,omitempty"`
	RoleSid                  *string    `json:"role_sid,omitempty"`
	ServiceSid               string     `json:"service_sid"`
	Sid                      string     `json:"sid"`
	URL                      string     `json:"url"`
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

// ChannelMembersPageResponse defines the response fields for the channel members page
type ChannelMembersPageResponse struct {
	Members []PageChannelMemberResponse `json:"members"`
	Meta    PageMetaResponse            `json:"meta"`
}

// Page retrieves a page of channel members
// See https://www.twilio.com/docs/chat/rest/member-resource#read-multiple-member-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ChannelMembersPageOptions) (*ChannelMembersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of channel members
// See https://www.twilio.com/docs/chat/rest/member-resource#read-multiple-member-resources for more details
func (c Client) PageWithContext(context context.Context, options *ChannelMembersPageOptions) (*ChannelMembersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Members",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChannelMembersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChannelMembersPaginator defines the fields for makings paginated api calls
// Members is an array of members that have been returned from all of the page calls
type ChannelMembersPaginator struct {
	options *ChannelMembersPageOptions
	Page    *ChannelMembersPage
	Members []PageChannelMemberResponse
}

// NewChannelMembersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChannelMembersPaginator() *ChannelMembersPaginator {
	return c.NewChannelMembersPaginatorWithOptions(nil)
}

// NewChannelMembersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChannelMembersPaginatorWithOptions(options *ChannelMembersPageOptions) *ChannelMembersPaginator {
	return &ChannelMembersPaginator{
		options: options,
		Page: &ChannelMembersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Members: make([]PageChannelMemberResponse, 0),
	}
}

// ChannelMembersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChannelMemberResponse or error that is returned from the api call(s)
type ChannelMembersPage struct {
	client *Client

	CurrentPage *ChannelMembersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChannelMembersPaginator) CurrentPage() *ChannelMembersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChannelMembersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChannelMembersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChannelMembersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChannelMembersPageOptions{}
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
		p.Members = append(p.Members, resp.Members...)
	}

	return p.Page.Error == nil
}
