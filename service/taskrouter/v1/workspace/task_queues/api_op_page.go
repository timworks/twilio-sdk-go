// Package task_queues contains auto-generated files. DO NOT MODIFY
package task_queues

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// TaskQueuesPageOptions defines the query options for the api operation
type TaskQueuesPageOptions struct {
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

type PageTaskQueueResponse struct {
	AccountSid              string     `json:"account_sid"`
	AssignmentActivityName  *string    `json:"assignment_activity_name,omitempty"`
	AssignmentActivitySid   *string    `json:"assignment_activity_sid,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	EventCallbackURL        *string    `json:"event_callback_url,omitempty"`
	FriendlyName            string     `json:"friendly_name"`
	MaxReservedWorkers      int        `json:"max_reserved_workers"`
	ReservationActivityName *string    `json:"reservation_activity_name,omitempty"`
	ReservationActivitySid  *string    `json:"reservation_activity_sid,omitempty"`
	Sid                     string     `json:"sid"`
	TargetWorkers           *string    `json:"target_workers,omitempty"`
	TaskOrder               string     `json:"task_order"`
	URL                     string     `json:"url"`
	WorkspaceSid            string     `json:"workspace_sid"`
}

// TaskQueuesPageResponse defines the response fields for the task queues page
type TaskQueuesPageResponse struct {
	Meta       PageMetaResponse        `json:"meta"`
	TaskQueues []PageTaskQueueResponse `json:"task_queues"`
}

// Page retrieves a page of task queues
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-list for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *TaskQueuesPageOptions) (*TaskQueuesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of task queues
// See https://www.twilio.com/docs/taskrouter/api/task-queue#action-list for more details
func (c Client) PageWithContext(context context.Context, options *TaskQueuesPageOptions) (*TaskQueuesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/TaskQueues",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &TaskQueuesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// TaskQueuesPaginator defines the fields for makings paginated api calls
// TaskQueues is an array of taskqueues that have been returned from all of the page calls
type TaskQueuesPaginator struct {
	options    *TaskQueuesPageOptions
	Page       *TaskQueuesPage
	TaskQueues []PageTaskQueueResponse
}

// NewTaskQueuesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewTaskQueuesPaginator() *TaskQueuesPaginator {
	return c.NewTaskQueuesPaginatorWithOptions(nil)
}

// NewTaskQueuesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewTaskQueuesPaginatorWithOptions(options *TaskQueuesPageOptions) *TaskQueuesPaginator {
	return &TaskQueuesPaginator{
		options: options,
		Page: &TaskQueuesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		TaskQueues: make([]PageTaskQueueResponse, 0),
	}
}

// TaskQueuesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageTaskQueueResponse or error that is returned from the api call(s)
type TaskQueuesPage struct {
	client *Client

	CurrentPage *TaskQueuesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *TaskQueuesPaginator) CurrentPage() *TaskQueuesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *TaskQueuesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *TaskQueuesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *TaskQueuesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &TaskQueuesPageOptions{}
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
		p.TaskQueues = append(p.TaskQueues, resp.TaskQueues...)
	}

	return p.Page.Error == nil
}
