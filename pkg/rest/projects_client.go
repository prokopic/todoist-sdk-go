package rest

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const (
	ProjectsPath = "projects"
)

// ProjectsClient is a REST client for working with Project resources.
type ProjectsClient struct {
	resourceClient resourceClient[Project]
}

// NewProjectsClient returns new REST client for projects.
func NewProjectsClient(httpClient *http.Client) (ProjectsClient, error) {
	if httpClient == nil {
		return ProjectsClient{}, errors.New("httpClient is not specified")
	}

	client := ProjectsClient{
		resourceClient: resourceClient[Project]{
			httpClient: httpClient,
			baseUrl:    APIUrl + "/" + ProjectsPath,
		},
	}

	return client, nil
}

// GetAll returns all projects.
func (c *ProjectsClient) GetAll(ctx context.Context) ([]Project, error) {
	return c.resourceClient.GetAll(ctx)
}

// Get returns a Project with specified ID.
func (c *ProjectsClient) Get(ctx context.Context, id int64) (Project, error) {
	return c.resourceClient.Get(ctx, id)
}
