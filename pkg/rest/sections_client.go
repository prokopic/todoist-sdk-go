package rest

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const (
	SectionsPath = "sections"
)

// SectionsClient is a REST client for working with Section resources.
type SectionsClient struct {
	resourceClient resourceClient[Section]
}

// NewSectionsClient returns new REST client for sections.
func NewSectionsClient(httpClient *http.Client) (SectionsClient, error) {
	if httpClient == nil {
		return SectionsClient{}, errors.New("httpClient is not specified")
	}

	client := SectionsClient{
		resourceClient: resourceClient[Section]{
			httpClient: httpClient,
			baseUrl:    APIUrl + "/" + SectionsPath,
		},
	}

	return client, nil
}

// GetForProject returns all project sections.
func (c *SectionsClient) GetForProject(ctx context.Context, projectID int64) ([]Section, error) {
	return c.resourceClient.GetForParent(ctx, "project_id", projectID)
}

// Get returns a Section with specified ID.
func (c *SectionsClient) Get(ctx context.Context, id int64) (Section, error) {
	return c.resourceClient.Get(ctx, id)
}
