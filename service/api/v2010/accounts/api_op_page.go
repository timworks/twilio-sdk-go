// Package accounts contains auto-generated files. DO NOT MODIFY
package accounts

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// AccountsPageOptions defines the query options for the api operation
type AccountsPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	Status       *string
}

type PageAccountResponse struct {
	AuthToken       string             `json:"auth_token"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	OwnerAccountSid string             `json:"owner_account_sid"`
	Sid             string             `json:"sid"`
	Status          string             `json:"status"`
	Type            string             `json:"type"`
}

// AccountsPageResponse defines the response fields for the accounts page
type AccountsPageResponse struct {
	Accounts        []PageAccountResponse `json:"accounts"`
	End             int                   `json:"end"`
	FirstPageURI    string                `json:"first_page_uri"`
	NextPageURI     *string               `json:"next_page_uri,omitempty"`
	Page            int                   `json:"page"`
	PageSize        int                   `json:"page_size"`
	PreviousPageURI *string               `json:"previous_page_uri,omitempty"`
	Start           int                   `json:"start"`
	URI             string                `json:"uri"`
}

// Page retrieves a page of account
// See https://www.twilio.com/docs/iam/api/account#read-multiple-account-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *AccountsPageOptions) (*AccountsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of account
// See https://www.twilio.com/docs/iam/api/account#read-multiple-account-resources for more details
func (c Client) PageWithContext(context context.Context, options *AccountsPageOptions) (*AccountsPageResponse, error) {
	op := client.Operation{
		Method:      http.MethodGet,
		URI:         "/Accounts.json",
		QueryParams: utils.StructToURLValues(options),
	}

	response := &AccountsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AccountsPaginator defines the fields for makings paginated api calls
// Accounts is an array of accounts that have been returned from all of the page calls
type AccountsPaginator struct {
	options  *AccountsPageOptions
	Page     *AccountsPage
	Accounts []PageAccountResponse
}

// NewAccountsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewAccountsPaginator() *AccountsPaginator {
	return c.NewAccountsPaginatorWithOptions(nil)
}

// NewAccountsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewAccountsPaginatorWithOptions(options *AccountsPageOptions) *AccountsPaginator {
	return &AccountsPaginator{
		options: options,
		Page: &AccountsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Accounts: make([]PageAccountResponse, 0),
	}
}

// AccountsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageAccountResponse or error that is returned from the api call(s)
type AccountsPage struct {
	client *Client

	CurrentPage *AccountsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *AccountsPaginator) CurrentPage() *AccountsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *AccountsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *AccountsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *AccountsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &AccountsPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPage := p.CurrentPage().NextPageURI

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
		p.Accounts = append(p.Accounts, resp.Accounts...)
	}

	return p.Page.Error == nil
}
