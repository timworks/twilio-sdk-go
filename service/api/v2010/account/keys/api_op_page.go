// Package keys contains auto-generated files. DO NOT MODIFY
package keys

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// KeysPageOptions defines the query options for the api operation
type KeysPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
	Status       *string
}

type PageKeyResponse struct {
	DateCreated  utils.RFC2822Time  `json:"date_created"`
	DateUpdated  *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName string             `json:"friendly_name"`
	Sid          string             `json:"sid"`
}

// KeysPageResponse defines the response fields for the keys page
type KeysPageResponse struct {
	End             int               `json:"end"`
	FirstPageURI    string            `json:"first_page_uri"`
	Keys            []PageKeyResponse `json:"keys"`
	NextPageURI     *string           `json:"next_page_uri,omitempty"`
	Page            int               `json:"page"`
	PageSize        int               `json:"page_size"`
	PreviousPageURI *string           `json:"previous_page_uri,omitempty"`
	Start           int               `json:"start"`
	URI             string            `json:"uri"`
}

// Page retrieves a page of keys
// See https://www.twilio.com/docs/iam/keys/api-key-resource#read-a-key-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *KeysPageOptions) (*KeysPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of keys
// See https://www.twilio.com/docs/iam/keys/api-key-resource#read-a-key-resource for more details
func (c Client) PageWithContext(context context.Context, options *KeysPageOptions) (*KeysPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Keys.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &KeysPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// KeysPaginator defines the fields for makings paginated api calls
// Keys is an array of keys that have been returned from all of the page calls
type KeysPaginator struct {
	options *KeysPageOptions
	Page    *KeysPage
	Keys    []PageKeyResponse
}

// NewKeysPaginator creates a new instance of the paginator for Page.
func (c *Client) NewKeysPaginator() *KeysPaginator {
	return c.NewKeysPaginatorWithOptions(nil)
}

// NewKeysPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewKeysPaginatorWithOptions(options *KeysPageOptions) *KeysPaginator {
	return &KeysPaginator{
		options: options,
		Page: &KeysPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Keys: make([]PageKeyResponse, 0),
	}
}

// KeysPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageKeyResponse or error that is returned from the api call(s)
type KeysPage struct {
	client *Client

	CurrentPage *KeysPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *KeysPaginator) CurrentPage() *KeysPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *KeysPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *KeysPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *KeysPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &KeysPageOptions{}
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
		p.Keys = append(p.Keys, resp.Keys...)
	}

	return p.Page.Error == nil
}
