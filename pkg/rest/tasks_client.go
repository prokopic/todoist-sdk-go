package rest

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const (
	TasksPath = "tasks"
)

// TasksClient is a REST client for working with Task resources.
type TasksClient struct {
	resourceClient resourceClient[Task]
}

// NewTasksClient returns new REST client for tasks.
func NewTasksClient(httpClient *http.Client) (TasksClient, error) {
	if httpClient == nil {
		return TasksClient{}, errors.New("httpClient is not specified")
	}

	client := TasksClient{
		resourceClient: resourceClient[Task]{
			httpClient: httpClient,
			baseUrl:    APIUrl + "/" + TasksPath,
		},
	}

	return client, nil
}

// GetForProject returns all project tasks.
func (c *TasksClient) GetForProject(ctx context.Context, projectID int64) ([]Task, error) {
	return c.resourceClient.GetForParent(ctx, "project_id", projectID)
}

// GetForSection returns all tasks from specified project section.
func (c *TasksClient) GetForSection(ctx context.Context, sectionID int64) ([]Task, error) {
	return c.resourceClient.GetForParent(ctx, "section_id", sectionID)
}

// GetWithLabel returns all tasks with specified label.
func (c *TasksClient) GetWithLabel(ctx context.Context, labelID int64) ([]Task, error) {
	return c.resourceClient.GetForParent(ctx, "label_id", labelID)
}

// GetWithFilter returns all tasks that satisfy specified filter.
func (c *TasksClient) GetWithFilter(ctx context.Context, filter string) ([]Task, error) {
	return c.resourceClient.GetWithParams(ctx, map[string]string{"filter": filter})
}

// Get returns a Task with specified ID.
func (c *TasksClient) Get(ctx context.Context, id int64) (Task, error) {
	return c.resourceClient.Get(ctx, id)
}
