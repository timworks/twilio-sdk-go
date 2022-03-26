// Package real_time_statistics contains auto-generated files. DO NOT MODIFY
package real_time_statistics

import (
	"context"
	"net/http"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchRealTimeStatisticsOptions defines the query options for the api operation
type FetchRealTimeStatisticsOptions struct {
	TaskChannel *string
}

type FetchActivityStatistic struct {
	FriendlyName string `json:"friendly_name"`
	Sid          string `json:"sid"`
	Workers      int    `json:"workers"`
}

// FetchRealTimeStatisticsResponse defines the response fields for the retrieved real time statistics
type FetchRealTimeStatisticsResponse struct {
	AccountSid         string                   `json:"account_sid"`
	ActivityStatistics []FetchActivityStatistic `json:"activity_statistics"`
	TotalWorkers       int                      `json:"total_workers"`
	URL                string                   `json:"url"`
	WorkspaceSid       string                   `json:"workspace_sid"`
}

// Fetch retrieves real time statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-real-time-worker-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchRealTimeStatisticsOptions) (*FetchRealTimeStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves real time statistics
// See https://www.twilio.com/docs/taskrouter/api/worker/statistics#fetch-real-time-worker-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchRealTimeStatisticsOptions) (*FetchRealTimeStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workers/RealTimeStatistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchRealTimeStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
