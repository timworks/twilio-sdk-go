// Package workspace contains auto-generated files. DO NOT MODIFY
package workspace

import (
	"github.com/timworks/twilio-sdk-go/client"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/activities"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/activity"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/cumulative_statistics"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/real_time_statistics"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/statistics"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/task"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/task_channel"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/task_channels"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/tasks"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/worker"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workflow"
	"github.com/timworks/twilio-sdk-go/service/taskrouter/v1/workspace/workflows"
)

// Client for managing a specific workspace resource
// See https://www.twilio.com/docs/taskrouter/api/workspace for more details
type Client struct {
	client *client.Client

	sid string

	Activities           *activities.Client
	Activity             func(string) *activity.Client
	CumulativeStatistics func() *cumulative_statistics.Client
	RealTimeStatistics   func() *real_time_statistics.Client
	Statistics           func() *statistics.Client
	Task                 func(string) *task.Client
	TaskChannel          func(string) *task_channel.Client
	TaskChannels         *task_channels.Client
	TaskQueue            func(string) *task_queue.Client
	TaskQueues           *task_queues.Client
	Tasks                *tasks.Client
	Worker               func(string) *worker.Client
	Workers              *workers.Client
	Workflow             func(string) *workflow.Client
	Workflows            *workflows.Client
}

// ClientProperties are the properties required to manage the workspace resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the workspace client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Activities: activities.New(client, activities.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
		Activity: func(activitySid string) *activity.Client {
			return activity.New(client, activity.ClientProperties{
				Sid:          activitySid,
				WorkspaceSid: properties.Sid,
			})
		},
		CumulativeStatistics: func() *cumulative_statistics.Client {
			return cumulative_statistics.New(client, cumulative_statistics.ClientProperties{
				WorkspaceSid: properties.Sid,
			})
		},
		RealTimeStatistics: func() *real_time_statistics.Client {
			return real_time_statistics.New(client, real_time_statistics.ClientProperties{
				WorkspaceSid: properties.Sid,
			})
		},
		Statistics: func() *statistics.Client {
			return statistics.New(client, statistics.ClientProperties{
				WorkspaceSid: properties.Sid,
			})
		},
		Task: func(taskSid string) *task.Client {
			return task.New(client, task.ClientProperties{
				Sid:          taskSid,
				WorkspaceSid: properties.Sid,
			})
		},
		TaskChannel: func(taskChannelSid string) *task_channel.Client {
			return task_channel.New(client, task_channel.ClientProperties{
				Sid:          taskChannelSid,
				WorkspaceSid: properties.Sid,
			})
		},
		TaskChannels: task_channels.New(client, task_channels.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
		TaskQueue: func(taskQueueSid string) *task_queue.Client {
			return task_queue.New(client, task_queue.ClientProperties{
				Sid:          taskQueueSid,
				WorkspaceSid: properties.Sid,
			})
		},
		TaskQueues: task_queues.New(client, task_queues.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
		Tasks: tasks.New(client, tasks.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
		Worker: func(workerSid string) *worker.Client {
			return worker.New(client, worker.ClientProperties{
				Sid:          workerSid,
				WorkspaceSid: properties.Sid,
			})
		},
		Workers: workers.New(client, workers.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
		Workflow: func(workflowSid string) *workflow.Client {
			return workflow.New(client, workflow.ClientProperties{
				Sid:          workflowSid,
				WorkspaceSid: properties.Sid,
			})
		},
		Workflows: workflows.New(client, workflows.ClientProperties{
			WorkspaceSid: properties.Sid,
		}),
	}
}
