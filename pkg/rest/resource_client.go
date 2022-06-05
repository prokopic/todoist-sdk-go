package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type resource interface {
	Comment | Label | Project | Section | Task
}

type resourceClientConfig struct {
	httpClient   *http.Client
	resourcePath string
}

type resourceClient[T resource] struct {
	httpClient *http.Client
	baseUrl    string
}

func newResourceClient[T resource](config resourceClientConfig) (resourceClient[T], error) {
	if config.httpClient == nil {
		return resourceClient[T]{}, errors.New("config.httpClient is not specified")
	}
	if config.resourcePath == "" {
		return resourceClient[T]{}, errors.New("config.resourcePath is not specified")
	}

	rc := resourceClient[T]{
		httpClient: config.httpClient,
		baseUrl:    APIUrl + "/" + config.resourcePath,
	}

	return rc, nil
}

// GetAll returns all resources of type T
func (c *resourceClient[T]) GetAll(ctx context.Context) ([]T, error) {
	return get[[]T](ctx, c.httpClient, c.baseUrl, nil)
}

// GetForParent returns all resources of type T
func (c *resourceClient[T]) GetForParent(ctx context.Context, parentIdName string, parentID int64) ([]T, error) {
	return get[[]T](ctx, c.httpClient, c.baseUrl, map[string]string{parentIdName: strconv.FormatInt(parentID, 10)})
}

// GetWithParams returns all resources of type T with specified query parameters
func (c *resourceClient[T]) GetWithParams(ctx context.Context, query map[string]string) ([]T, error) {
	return get[[]T](ctx, c.httpClient, c.baseUrl, query)
}

func (c *resourceClient[T]) Get(ctx context.Context, id int64) (T, error) {
	return get[T](ctx, c.httpClient, c.baseUrl+"/"+strconv.FormatInt(id, 10), nil)
}

func get[T any](ctx context.Context, client *http.Client, url string, query map[string]string) (T, error) {
	token := os.Getenv(APITokenEnvVar)
	if token == "" {
		return *new(T), errors.New(fmt.Sprintf("failed to get Todoist API token, env var %s is not set", APITokenEnvVar))
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return *new(T), errors.Wrapf(err, "failed to create new GET request (GET %s)", url)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	q := request.URL.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		return *new(T), errors.Wrapf(err, "failed to complete GET request (GET %s)", url)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return *new(T), errors.Wrapf(err, "failed to read from response body (GET %s)", url)
	}

	var value T
	err = json.Unmarshal(body, &value)
	if err != nil {
		return *new(T), errors.Wrapf(err, "failed to unmarshal value from response body (GET %s): %s", url, body)
	}

	return value, nil
}
