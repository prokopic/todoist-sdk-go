package rest

import (
	"github.com/pkg/errors"
	"net/http"
)

const (
	APIBaseUrl = "https://api.todoist.com/rest"
	APIVersion = "v1"
	APIUrl     = APIBaseUrl + "/" + APIVersion

	APITokenEnvVar = "TODOIST_API_TOKEN"
)

type Client struct {
	Projects ResourceClient[Project]
	Labels   ResourceClient[Label]
}

type resource interface {
	Comment | Label | Project | Section | Task
}

func NewClient(httpClient *http.Client) (*Client, error) {
	projectsClientConfig := resourceClientConfig{
		httpClient: httpClient,
		baseUrl:    ProjectsUrl,
	}
	projectsClient, err := newResourceClient[Project](projectsClientConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create Projects client")
	}

	labelsClientConfig := resourceClientConfig{
		httpClient: httpClient,
		baseUrl:    LabelsUrl,
	}
	labelsClient, err := newResourceClient[Label](labelsClientConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create Labels client")
	}

	client := Client{
		Projects: projectsClient,
		Labels:   labelsClient,
	}
	return &client, nil
}
