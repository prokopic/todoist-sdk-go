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
	Projects ProjectsClient
	Sections SectionsClient
	Tasks    TasksClient
	Comments CommentsClient
	Labels   LabelsClient
}

func NewClient(httpClient *http.Client) (Client, error) {
	projectsClient, err := NewProjectsClient(httpClient)
	if err != nil {
		return Client{}, errors.WithMessage(err, "failed to create Projects client")
	}

	sectionsClient, err := NewSectionsClient(httpClient)
	if err != nil {
		return Client{}, errors.WithMessage(err, "failed to create Sections client")
	}

	tasksClient, err := NewTasksClient(httpClient)
	if err != nil {
		return Client{}, errors.WithMessage(err, "failed to create Tasks client")
	}

	commentsClient, err := NewCommentsClient(httpClient)
	if err != nil {
		return Client{}, errors.WithMessage(err, "failed to create Comments client")
	}

	labelsClient, err := NewLabelsClient(httpClient)
	if err != nil {
		return Client{}, errors.WithMessage(err, "failed to create Labels client")
	}

	client := Client{
		Projects: projectsClient,
		Sections: sectionsClient,
		Tasks:    tasksClient,
		Comments: commentsClient,
		Labels:   labelsClient,
	}

	return client, nil
}
