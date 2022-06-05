package rest

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const (
	CommentsPath = "comments"
)

// CommentsClient is a REST client for working with Comment resources.
type CommentsClient struct {
	resourceClient resourceClient[Comment]
}

// NewCommentsClient returns new REST client for comments.
func NewCommentsClient(httpClient *http.Client) (CommentsClient, error) {
	if httpClient == nil {
		return CommentsClient{}, errors.New("httpClient is not specified")
	}

	client := CommentsClient{
		resourceClient: resourceClient[Comment]{
			httpClient: httpClient,
			baseUrl:    APIUrl + "/" + CommentsPath,
		},
	}

	return client, nil
}

// GetForProject returns all project comments.
func (c *CommentsClient) GetForProject(ctx context.Context, projectID int64) ([]Comment, error) {
	return c.resourceClient.GetForParent(ctx, "project_id", projectID)
}

// GetForTask returns all task comments.
func (c *CommentsClient) GetForTask(ctx context.Context, taskID int64) ([]Comment, error) {
	return c.resourceClient.GetForParent(ctx, "task_id", taskID)
}

// Get returns a Comment with specified ID.
func (c *CommentsClient) Get(ctx context.Context, id int64) (Comment, error) {
	return c.resourceClient.Get(ctx, id)
}
