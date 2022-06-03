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

type ResourceClient[T resource] interface {
	GetAll(context.Context) ([]T, error)
	Get(context.Context, int64) (T, error)
}

type resourceClientConfig struct {
	httpClient *http.Client
	baseUrl    string
}

type resourceClientImpl[T resource] struct {
	httpClient *http.Client
	baseUrl    string
}

func newResourceClient[T resource](config resourceClientConfig) (*resourceClientImpl[T], error) {
	if config.httpClient == nil {
		return nil, errors.New("config.httpClient is not specified")
	}
	if config.baseUrl == "" {
		return nil, errors.New("config.baseUrl is not specified")
	}

	rc := resourceClientImpl[T]{
		httpClient: config.httpClient,
		baseUrl:    config.baseUrl,
	}

	return &rc, nil
}

func (c *resourceClientImpl[T]) GetAll(ctx context.Context) ([]T, error) {
	return get[[]T](ctx, c.httpClient, c.baseUrl)
}

func (c *resourceClientImpl[T]) Get(ctx context.Context, id int64) (T, error) {
	return get[T](ctx, c.httpClient, c.baseUrl+"/"+strconv.FormatInt(id, 10))
}

func get[T any](ctx context.Context, client *http.Client, url string) (T, error) {
	token := os.Getenv(APITokenEnvVar)
	if token == "" {
		return *new(T), errors.New(fmt.Sprintf("failed to get Todoist API token, env var %s is not set", APITokenEnvVar))
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return *new(T), errors.Wrapf(err, "failed to create new GET request (GET %s)", url)
	}
	request.Header.Set("Authorization", "Bearer "+token)

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
