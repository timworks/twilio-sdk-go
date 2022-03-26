// Package workers contains auto-generated files. DO NOT MODIFY
package workers

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// WorkersPageOptions defines the query options for the api operation
type WorkersPageOptions struct {
	PageSize                *int
	Page                    *int
	PageToken               *string
	ActivityName            *string
	ActivitySid             *string
	Available               *bool
	FriendlyName            *string
	TargetWorkersExpression *string
	TaskQueueName           *string
	TaskQueueSid            *string
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

type PageWorkerResponse struct {
	AccountSid        string     `json:"account_sid"`
	ActivityName      string     `json:"activity_name"`
	ActivitySid       string     `json:"activity_sid"`
	Attributes        string     `json:"attributes"`
	Available         bool       `json:"available"`
	DateCreated       time.Time  `json:"date_created"`
	DateStatusChanged *time.Time `json:"date_status_changed,omitempty"`
	DateUpdated       *time.Time `json:"date_updated,omitempty"`
	FriendlyName      string     `json:"friendly_name"`
	Sid               string     `json:"sid"`
	URL               string     `json:"url"`
	WorkspaceSid      string     `json:"workspace_sid"`
}

// WorkersPageResponse defines the response fields for the workers page
type WorkersPageResponse struct {
	Meta    PageMetaResponse     `json:"meta"`
	Workers []PageWorkerResponse `json:"workers"`
}

// Page retrieves a page of workers
// See https://www.twilio.com/docs/taskrouter/api/worker#read-multiple-worker-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *WorkersPageOptions) (*WorkersPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of workers
// See https://www.twilio.com/docs/taskrouter/api/worker#read-multiple-worker-resources for more details
func (c Client) PageWithContext(context context.Context, options *WorkersPageOptions) (*WorkersPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &WorkersPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// WorkersPaginator defines the fields for makings paginated api calls
// Workers is an array of workers that have been returned from all of the page calls
type WorkersPaginator struct {
	options *WorkersPageOptions
	Page    *WorkersPage
	Workers []PageWorkerResponse
}

// NewWorkersPaginator creates a new instance of the paginator for Page.
func (c *Client) NewWorkersPaginator() *WorkersPaginator {
	return c.NewWorkersPaginatorWithOptions(nil)
}

// NewWorkersPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewWorkersPaginatorWithOptions(options *WorkersPageOptions) *WorkersPaginator {
	return &WorkersPaginator{
		options: options,
		Page: &WorkersPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Workers: make([]PageWorkerResponse, 0),
	}
}

// WorkersPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageWorkerResponse or error that is returned from the api call(s)
type WorkersPage struct {
	client *Client

	CurrentPage *WorkersPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *WorkersPaginator) CurrentPage() *WorkersPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *WorkersPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *WorkersPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *WorkersPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &WorkersPageOptions{}
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
		p.Workers = append(p.Workers, resp.Workers...)
	}

	return p.Page.Error == nil
}
