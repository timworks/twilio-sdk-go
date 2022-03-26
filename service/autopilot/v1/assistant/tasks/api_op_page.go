// Package tasks contains auto-generated files. DO NOT MODIFY
package tasks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// TasksPageOptions defines the query options for the api operation
type TasksPageOptions struct {
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

type PageTaskResponse struct {
	AccountSid   string     `json:"account_sid"`
	ActionsURL   string     `json:"actions_url"`
	AssistantSid string     `json:"assistant_sid"`
	DateCreated  time.Time  `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated,omitempty"`
	FriendlyName *string    `json:"friendly_name,omitempty"`
	Sid          string     `json:"sid"`
	URL          string     `json:"url"`
	UniqueName   string     `json:"unique_name"`
}

// TasksPageResponse defines the response fields for the tasks page
type TasksPageResponse struct {
	Meta  PageMetaResponse   `json:"meta"`
	Tasks []PageTaskResponse `json:"tasks"`
}

// Page retrieves a page of tasks
// See https://www.twilio.com/docs/autopilot/api/task#read-multiple-task-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *TasksPageOptions) (*TasksPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of tasks
// See https://www.twilio.com/docs/autopilot/api/task#read-multiple-task-resources for more details
func (c Client) PageWithContext(context context.Context, options *TasksPageOptions) (*TasksPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Assistants/{assistantSid}/Tasks",
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &TasksPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// TasksPaginator defines the fields for makings paginated api calls
// Tasks is an array of tasks that have been returned from all of the page calls
type TasksPaginator struct {
	options *TasksPageOptions
	Page    *TasksPage
	Tasks   []PageTaskResponse
}

// NewTasksPaginator creates a new instance of the paginator for Page.
func (c *Client) NewTasksPaginator() *TasksPaginator {
	return c.NewTasksPaginatorWithOptions(nil)
}

// NewTasksPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewTasksPaginatorWithOptions(options *TasksPageOptions) *TasksPaginator {
	return &TasksPaginator{
		options: options,
		Page: &TasksPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Tasks: make([]PageTaskResponse, 0),
	}
}

// TasksPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageTaskResponse or error that is returned from the api call(s)
type TasksPage struct {
	client *Client

	CurrentPage *TasksPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *TasksPaginator) CurrentPage() *TasksPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *TasksPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *TasksPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *TasksPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &TasksPageOptions{}
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
		p.Tasks = append(p.Tasks, resp.Tasks...)
	}

	return p.Page.Error == nil
}
