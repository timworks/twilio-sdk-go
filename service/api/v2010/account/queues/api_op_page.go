// Package queues contains auto-generated files. DO NOT MODIFY
package queues

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// QueuesPageOptions defines the query options for the api operation
type QueuesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
}

type PageQueueResponse struct {
	AccountSid      string             `json:"account_sid"`
	AverageWaitTime int                `json:"average_wait_time"`
	CurrentSize     int                `json:"current_size"`
	DateCreated     utils.RFC2822Time  `json:"date_created"`
	DateUpdated     *utils.RFC2822Time `json:"date_updated,omitempty"`
	FriendlyName    string             `json:"friendly_name"`
	MaxSize         int                `json:"max_size"`
	Sid             string             `json:"sid"`
}

// QueuesPageResponse defines the response fields for the queues page
type QueuesPageResponse struct {
	End             int                 `json:"end"`
	FirstPageURI    string              `json:"first_page_uri"`
	NextPageURI     *string             `json:"next_page_uri,omitempty"`
	Page            int                 `json:"page"`
	PageSize        int                 `json:"page_size"`
	PreviousPageURI *string             `json:"previous_page_uri,omitempty"`
	Queues          []PageQueueResponse `json:"queues"`
	Start           int                 `json:"start"`
	URI             string              `json:"uri"`
}

// Page retrieves a page of queues
// See https://www.twilio.com/docs/voice/api/queue-resource#read-multiple-queue-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *QueuesPageOptions) (*QueuesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of queues
// See https://www.twilio.com/docs/voice/api/queue-resource#read-multiple-queue-resources for more details
func (c Client) PageWithContext(context context.Context, options *QueuesPageOptions) (*QueuesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Queues.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &QueuesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// QueuesPaginator defines the fields for makings paginated api calls
// Queues is an array of queues that have been returned from all of the page calls
type QueuesPaginator struct {
	options *QueuesPageOptions
	Page    *QueuesPage
	Queues  []PageQueueResponse
}

// NewQueuesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewQueuesPaginator() *QueuesPaginator {
	return c.NewQueuesPaginatorWithOptions(nil)
}

// NewQueuesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewQueuesPaginatorWithOptions(options *QueuesPageOptions) *QueuesPaginator {
	return &QueuesPaginator{
		options: options,
		Page: &QueuesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Queues: make([]PageQueueResponse, 0),
	}
}

// QueuesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageQueueResponse or error that is returned from the api call(s)
type QueuesPage struct {
	client *Client

	CurrentPage *QueuesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *QueuesPaginator) CurrentPage() *QueuesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *QueuesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *QueuesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *QueuesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &QueuesPageOptions{}
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
		p.Queues = append(p.Queues, resp.Queues...)
	}

	return p.Page.Error == nil
}
