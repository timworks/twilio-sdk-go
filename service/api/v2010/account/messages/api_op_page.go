// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// MessagesPageOptions defines the query options for the api operation
type MessagesPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
	To        *string
	From      *string
	DateSent  *string
}

type PageMessageResponse struct {
	APIVersion          string             `json:"api_version"`
	AccountSid          string             `json:"account_sid"`
	Body                string             `json:"body"`
	DateCreated         utils.RFC2822Time  `json:"date_created"`
	DateSent            utils.RFC2822Time  `json:"date_sent"`
	DateUpdated         *utils.RFC2822Time `json:"date_updated,omitempty"`
	Direction           string             `json:"direction"`
	ErrorCode           *int               `json:"error_code,omitempty"`
	ErrorMessage        *string            `json:"error_message,omitempty"`
	From                *string            `json:"from,omitempty"`
	MessagingServiceSid *string            `json:"messaging_service_sid,omitempty"`
	NumMedia            string             `json:"num_media"`
	NumSegments         string             `json:"num_segments"`
	Price               *string            `json:"price,omitempty"`
	PriceUnit           string             `json:"price_unit"`
	Sid                 string             `json:"sid"`
	Status              string             `json:"status"`
	To                  string             `json:"to"`
}

// MessagesPageResponse defines the response fields for the messages page
type MessagesPageResponse struct {
	End             int                   `json:"end"`
	FirstPageURI    string                `json:"first_page_uri"`
	Messages        []PageMessageResponse `json:"messages"`
	NextPageURI     *string               `json:"next_page_uri,omitempty"`
	Page            int                   `json:"page"`
	PageSize        int                   `json:"page_size"`
	PreviousPageURI *string               `json:"previous_page_uri,omitempty"`
	Start           int                   `json:"start"`
	URI             string                `json:"uri"`
}

// Page retrieves a page of messages
// See https://www.twilio.com/docs/sms/api/message-resource#read-multiple-message-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *MessagesPageOptions) (*MessagesPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of messages
// See https://www.twilio.com/docs/sms/api/message-resource#read-multiple-message-resources for more details
func (c Client) PageWithContext(context context.Context, options *MessagesPageOptions) (*MessagesPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Accounts/{accountSid}/Messages.json",
		PathParams: map[string]string{
			"accountSid": c.accountSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &MessagesPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// MessagesPaginator defines the fields for makings paginated api calls
// Messages is an array of messages that have been returned from all of the page calls
type MessagesPaginator struct {
	options  *MessagesPageOptions
	Page     *MessagesPage
	Messages []PageMessageResponse
}

// NewMessagesPaginator creates a new instance of the paginator for Page.
func (c *Client) NewMessagesPaginator() *MessagesPaginator {
	return c.NewMessagesPaginatorWithOptions(nil)
}

// NewMessagesPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewMessagesPaginatorWithOptions(options *MessagesPageOptions) *MessagesPaginator {
	return &MessagesPaginator{
		options: options,
		Page: &MessagesPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		Messages: make([]PageMessageResponse, 0),
	}
}

// MessagesPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageMessageResponse or error that is returned from the api call(s)
type MessagesPage struct {
	client *Client

	CurrentPage *MessagesPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *MessagesPaginator) CurrentPage() *MessagesPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *MessagesPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *MessagesPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *MessagesPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &MessagesPageOptions{}
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
		p.Messages = append(p.Messages, resp.Messages...)
	}

	return p.Page.Error == nil
}
