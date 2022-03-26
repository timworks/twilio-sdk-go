// Package invites contains auto-generated files. DO NOT MODIFY
package invites

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// ChannelInvitesPageOptions defines the query options for the api operation
type ChannelInvitesPageOptions struct {
	Identity  *[]string
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageChannelInviteResponse struct {
	AccountSid  string     `json:"account_sid"`
	ChannelSid  string     `json:"channel_sid"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    string     `json:"identity"`
	RoleSid     *string    `json:"role_sid,omitempty"`
	ServiceSid  string     `json:"service_sid"`
	Sid         string     `json:"sid"`
	URL         string     `json:"url"`
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

// ChannelInvitesPageResponse defines the response fields for the channel invites page
type ChannelInvitesPageResponse struct {
	Invites []PageChannelInviteResponse `json:"invites"`
	Meta    PageMetaResponse            `json:"meta"`
}

// Page retrieves a page of channel invites
// See https://www.twilio.com/docs/chat/rest/invite-resource#read-multiple-invite-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *ChannelInvitesPageOptions) (*ChannelInvitesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of channel invites
// See https://www.twilio.com/docs/chat/rest/invite-resource#read-multiple-invite-resources for more details
func (c Client) PageWithContext(context context.Context, options *ChannelInvitesPageOptions) (*ChannelInvitesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Channels/{channelSid}/Invites",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"channelSid": c.channelSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &ChannelInvitesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// ChannelInvitesPaginator defines the fields for makings paginated api calls
// Invites is an array of invites that have been returned from all of the page calls
type ChannelInvitesPaginator struct {
	options *ChannelInvitesPageOptions
	Page    *ChannelInvitesPage
	Invites []PageChannelInviteResponse
}

// NewChannelInvitesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewChannelInvitesPaginator() *ChannelInvitesPaginator {
	return c.NewChannelInvitesPaginatorWithOptions(nil)
}

// NewChannelInvitesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewChannelInvitesPaginatorWithOptions(options *ChannelInvitesPageOptions) *ChannelInvitesPaginator {
	return &ChannelInvitesPaginator{
		options: options,
		Page: &ChannelInvitesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Invites: make([]PageChannelInviteResponse, 0),
	}
}

// ChannelInvitesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageChannelInviteResponse or error that is returned from the api call(s)
type ChannelInvitesPage struct {
	client *Client

	CurrentPage *ChannelInvitesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *ChannelInvitesPaginator) CurrentPage() *ChannelInvitesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *ChannelInvitesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *ChannelInvitesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *ChannelInvitesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &ChannelInvitesPageOptions{}
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
		p.Invites = append(p.Invites, resp.Invites...)
	}

	return p.Page.Error == nil
}
