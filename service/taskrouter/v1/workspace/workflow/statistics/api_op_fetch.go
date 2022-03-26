// Package statistics contains auto-generated files. DO NOT MODIFY
package statistics

import (
	"context"
	"net/http"
	"time"

	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/utils"
)

// FetchStatisticsOptions defines the query options for the api operation
type FetchStatisticsOptions struct {
	Minutes         *int
	StartDate       *time.Time
	EndDate         *time.Time
	TaskChannel     *string
	SplitByWaitTime *string
}

type FetchCumulativeStatistics struct {
	AvgTaskAcceptanceTime     int                       `json:"avg_task_acceptance_time"`
	EndTime                   time.Time                 `json:"end_time"`
	ReservationsAccepted      int                       `json:"reservations_accepted"`
	ReservationsCanceled      int                       `json:"reservations_canceled"`
	ReservationsCompleted     *int                      `json:"reservations_completed,omitempty"`
	ReservationsCreated       int                       `json:"reservations_created"`
	ReservationsRejected      int                       `json:"reservations_rejected"`
	ReservationsRescinded     int                       `json:"reservations_rescinded"`
	ReservationsTimedOut      int                       `json:"reservations_timed_out"`
	ReservationsWrapUp        int                       `json:"reservations_wrapup"`
	SplitByWaitTime           *map[string]FetchWaitTime `json:"split_by_wait_time,omitempty"`
	StartTime                 time.Time                 `json:"start_time"`
	TasksAssigned             *int                      `json:"tasks_assigned,omitempty"`
	TasksCanceled             int                       `json:"tasks_canceled"`
	TasksCompleted            int                       `json:"tasks_completed"`
	TasksDeleted              int                       `json:"tasks_deleted"`
	TasksEntered              int                       `json:"tasks_entered"`
	TasksMoved                int                       `json:"tasks_moved"`
	TasksTimedOutInWorkflow   int                       `json:"tasks_timed_out_in_workflow"`
	WaitDurationUntilAccepted FetchStatisticsBreakdown  `json:"wait_duration_until_accepted"`
	WaitDurationUntilCanceled FetchStatisticsBreakdown  `json:"wait_duration_until_canceled"`
}

type FetchRealTimeStatistics struct {
	LongestTaskWaitingAge int            `json:"longest_task_waiting_age"`
	LongestTaskWaitingSid *string        `json:"longest_task_waiting_sid,omitempty"`
	TasksByPriority       map[string]int `json:"tasks_by_priority"`
	TasksByStatus         map[string]int `json:"tasks_by_status"`
	TotalTasks            int            `json:"total_tasks"`
}

type FetchStatisticsBreakdown struct {
	Avg   int `json:"avg"`
	Max   int `json:"max"`
	Min   int `json:"min"`
	Total int `json:"total"`
}

type FetchWaitTime struct {
	Above FetchWaitTimeTasks `json:"above"`
	Below FetchWaitTimeTasks `json:"below"`
}

type FetchWaitTimeTasks struct {
	ReservationsAccepted int `json:"reservations_accepted"`
	TasksCanceled        int `json:"tasks_canceled"`
}

// FetchStatisticsResponse defines the response fields for the retrieved statistics
type FetchStatisticsResponse struct {
	AccountSid   string                    `json:"account_sid"`
	Cumulative   FetchCumulativeStatistics `json:"cumulative"`
	RealTime     FetchRealTimeStatistics   `json:"realtime"`
	URL          string                    `json:"url"`
	WorkflowSid  string                    `json:"workflow_sid"`
	WorkspaceSid string                    `json:"workspace_sid"`
}

// Fetch retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/workflow-statistics#workflow-statistics for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Fetch(options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	return c.FetchWithContext(context.Background(), options)
}

// FetchWithContext retrieves statistics
// See https://www.twilio.com/docs/taskrouter/api/workflow-statistics#workflow-statistics for more details
func (c Client) FetchWithContext(context context.Context, options *FetchStatisticsOptions) (*FetchStatisticsResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Workspaces/{workspaceSid}/Workflows/{workflowSid}/Statistics",
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"workflowSid":  c.workflowSid,
		},
		QueryParams: utils.StructToURLValues(options),
	}

	response := &FetchStatisticsResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}
