// Package revisions contains auto-generated files. DO NOT MODIFY
package revisions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// RevisionsPageOptions defines the query options for the api operation
type RevisionsPageOptions struct {
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

type PageRevisionResponse struct {
	AccountSid    string                 `json:"account_sid"`
	CommitMessage *string                `json:"commit_message,omitempty"`
	DateCreated   time.Time              `json:"date_created"`
	DateUpdated   *time.Time             `json:"date_updated,omitempty"`
	Definition    map[string]interface{} `json:"definition"`
	Errors        *[]interface{}         `json:"errors,omitempty"`
	FriendlyName  string                 `json:"friendly_name"`
	Revision      int                    `json:"revision"`
	Sid           string                 `json:"sid"`
	Status        string                 `json:"status"`
	URL           string                 `json:"url"`
	Valid         bool                   `json:"valid"`
	Warnings      *[]interface{}         `json:"warnings,omitempty"`
	WebhookURL    string                 `json:"webhook_url"`
}

// RevisionsPageResponse defines the response fields for the flow revisions page
type RevisionsPageResponse struct {
	Meta      PageMetaResponse       `json:"meta"`
	Revisions []PageRevisionResponse `json:"revisions"`
}

// Page retrieves a page of flow revisions
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision#read-multiple-flowrevision-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *RevisionsPageOptions) (*RevisionsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of flow revisions
// See https://www.twilio.com/docs/studio/rest-api/v2/flow-revision#read-multiple-flowrevision-resources for more details
func (c Client) PageWithContext(context context.Context, options *RevisionsPageOptions) (*RevisionsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Flows/{flowSid}/Revisions",
		PathParams: map[string]string{
			"flowSid": c.flowSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &RevisionsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// RevisionsPaginator defines the fields for makings paginated api calls
// Revisions is an array of revisions that have been returned from all of the page calls
type RevisionsPaginator struct {
	options   *RevisionsPageOptions
	Page      *RevisionsPage
	Revisions []PageRevisionResponse
}

// NewRevisionsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewRevisionsPaginator() *RevisionsPaginator {
	return c.NewRevisionsPaginatorWithOptions(nil)
}

// NewRevisionsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewRevisionsPaginatorWithOptions(options *RevisionsPageOptions) *RevisionsPaginator {
	return &RevisionsPaginator{
		options: options,
		Page: &RevisionsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Revisions: make([]PageRevisionResponse, 0),
	}
}

// RevisionsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageRevisionResponse or error that is returned from the api call(s)
type RevisionsPage struct {
	client *Client

	CurrentPage *RevisionsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *RevisionsPaginator) CurrentPage() *RevisionsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *RevisionsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *RevisionsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *RevisionsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &RevisionsPageOptions{}
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
		p.Revisions = append(p.Revisions, resp.Revisions...)
	}

	return p.Page.Error == nil
}
