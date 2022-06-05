package rest

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
)

const (
	LabelsPath = "labels"
)

// LabelsClient is a REST client for working with Label resources.
type LabelsClient struct {
	resourceClient resourceClient[Label]
}

// NewLabelsClient returns new REST client for labels.
func NewLabelsClient(httpClient *http.Client) (LabelsClient, error) {
	if httpClient == nil {
		return LabelsClient{}, errors.New("httpClient is not specified")
	}

	client := LabelsClient{
		resourceClient: resourceClient[Label]{
			httpClient: httpClient,
			baseUrl:    APIUrl + "/" + LabelsPath,
		},
	}

	return client, nil
}

// GetAll returns all labels.
func (c *LabelsClient) GetAll(ctx context.Context) ([]Label, error) {
	return c.resourceClient.GetAll(ctx)
}

// Get returns a Label with specified ID.
func (c *LabelsClient) Get(ctx context.Context, id int64) (Label, error) {
	return c.resourceClient.Get(ctx, id)
}
