package peopledatalabs_go

import (
	"net/http"
	"time"

	"github.com/peopledatalabs/peopledatalabs-go/v6/api"
	"github.com/peopledatalabs/peopledatalabs-go/v6/logger"
)

// WithHTTPClient sets a custom HTTP Client
// that will be used to make API requests for the library
func WithHTTPClient(httpCli *http.Client) api.ClientOptions {
	return func(client *api.Client) {
		client.HttpClient = httpCli
	}
}

// WithTimeout sets the timeout the API client uses
// Default timeout is api.DefaultTimeout
func WithTimeout(timeout time.Duration) api.ClientOptions {
	return func(client *api.Client) {
		client.HttpClient.Timeout = timeout
	}
}

// WithLogger sets a custom logger for the library
func WithLogger(logger logger.Logger) api.ClientOptions {
	return func(client *api.Client) {
		client.Logger = logger
	}
}

// WithLogLevel sets the level at which the library will log
// Default behaviour is to log at logger.DefaultLevel
func WithLogLevel(level logger.Level) api.ClientOptions {
	return func(client *api.Client) {
		client.Logger.SetLevel(level)
	}
}
