// Package workflows contains auto-generated files. DO NOT MODIFY
package workflows

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// WorkflowsPageOptions defines the query options for the api operation
type WorkflowsPageOptions struct {
	PageSize     *int
	Page         *int
	PageToken    *string
	FriendlyName *string
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

type PageWorkflowResponse struct {
	AccountSid                    string     `json:"account_sid"`
	AssignmentCallbackURL         *string    `json:"assignment_callback_url,omitempty"`
	Configuration                 string     `json:"configuration"`
	DateCreated                   time.Time  `json:"date_created"`
	DateUpdated                   *time.Time `json:"date_updated,omitempty"`
	DocumentContentType           string     `json:"document_content_type"`
	FallbackAssignmentCallbackURL *string    `json:"fallback_assignment_callback_url,omitempty"`
	FriendlyName                  string     `json:"friendly_name"`
	Sid                           string     `json:"sid"`
	TaskReservationTimeout        int        `json:"task_reservation_timeout"`
	URL                           string     `json:"url"`
	WorkspaceSid                  string     `json:"workspace_sid"`
}

// WorkflowsPageResponse defines the response fields for the workflows page
type WorkflowsPageResponse struct {
	Meta      PageMetaResponse       `json:"meta"`
	Workflows []PageWorkflowResponse `json:"workflows"`
}

// Page retrieves a page of workflows
// See https://www.twilio.com/docs/taskrouter/api/workflow#read-multiple-workflow-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *WorkflowsPageOptions) (*WorkflowsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of workflows
// See https://www.twilio.com/docs/taskrouter/api/workflow#read-multiple-workflow-resources for more details
func (c Client) PageWithContext(context context.Context, options *WorkflowsPageOptions) (*WorkflowsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workflows",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &WorkflowsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// WorkflowsPaginator defines the fields for makings paginated api calls
// Workflows is an array of workflows that have been returned from all of the page calls
type WorkflowsPaginator struct {
	options   *WorkflowsPageOptions
	Page      *WorkflowsPage
	Workflows []PageWorkflowResponse
}

// NewWorkflowsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewWorkflowsPaginator() *WorkflowsPaginator {
	return c.NewWorkflowsPaginatorWithOptions(nil)
}

// NewWorkflowsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewWorkflowsPaginatorWithOptions(options *WorkflowsPageOptions) *WorkflowsPaginator {
	return &WorkflowsPaginator{
		options: options,
		Page: &WorkflowsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Workflows: make([]PageWorkflowResponse, 0),
	}
}

// WorkflowsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageWorkflowResponse or error that is returned from the api call(s)
type WorkflowsPage struct {
	client *Client

	CurrentPage *WorkflowsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *WorkflowsPaginator) CurrentPage() *WorkflowsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *WorkflowsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *WorkflowsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *WorkflowsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &WorkflowsPageOptions{}
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
		p.Workflows = append(p.Workflows, resp.Workflows...)
	}

	return p.Page.Error == nil
}
