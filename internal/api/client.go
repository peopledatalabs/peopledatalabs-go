package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/peopledatalabs/peopledatalabs-go/logger"
	"github.com/peopledatalabs/peopledatalabs-go/model"

	"github.com/google/go-querystring/query"
)

const (
	DefaultTimeout    = 10 * time.Second
	defaultBaseURL    = "https://api.peopledatalabs.com/"
	defaultApiVersion = "v5"
	defaultSandbox    = false
	defaultSandboxURL = "https://sandbox.api.peopledatalabs.com/"

)

type Client struct {
	ApiKey           string
	BaseURL          string
	ApiVersion       string
	HttpClient       *http.Client
	Logger           logger.Logger
	LibVersion       string
	Sandbox	         bool
	SandboxBaseURL   string
}

func NewClient(apiKey, libVersion string, opts ...ClientOptions) Client {
	c := Client{
		ApiKey:     apiKey,
		BaseURL:    defaultBaseURL,
		ApiVersion: defaultApiVersion,
		LibVersion: libVersion,
		HttpClient: defaultHTTPClient(),
		Logger:     logger.NewDefaultLogger(),
		Sandbox:    defaultSandbox,
		SandboxBaseURL: defaultSandboxURL,
	}

	if c.Sandbox {
		c.BaseURL = c.SandboxBaseURL
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: DefaultTimeout,
	}
}

type ClientOptions func(*Client)

type apiRequest struct {
	method  string
	path    string
	headers http.Header
	params  string
	body    []byte
}

func (cli Client) get(ctx context.Context, path string, params interface{}, out interface{}) error {
	queryParams, err := query.Values(params)
	if err != nil {
		return fmt.Errorf("makeRequest: could not generate params query: %w", err)
	}

	req := apiRequest{
		method: http.MethodGet,
		path:   path,
		params: queryParams.Encode(),
	}

	return cli.makeRequest(ctx, req, out)
}

func (cli Client) post(ctx context.Context, path string, params interface{}, out interface{}) error {
	body, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("makeRequest: could not marshal JSON: %w", err)
	}

	req := apiRequest{
		method:  http.MethodPost,
		path:    path,
		body:    body,
		headers: http.Header{},
	}
	req.headers.Set("Content-Type", "application/json")

	return cli.makeRequest(ctx, req, out)
}

func (cli Client) makeRequest(ctx context.Context, libReq apiRequest, out interface{}) error {
	urlStruct, err := url.Parse(fmt.Sprintf("%s%s%s", cli.BaseURL, cli.ApiVersion, libReq.path))
	if err != nil {
		return fmt.Errorf("makeRequest: %w", err)
	}
	urlStruct.RawQuery = libReq.params
	URL := urlStruct.String()

	cli.Logger.Debug(requestLog(URL, libReq))

	req, err := http.NewRequestWithContext(ctx, libReq.method, URL, bytes.NewReader(libReq.body))
	if err != nil {
		return err
	}

	userAgent := fmt.Sprintf("peopledatalabs-go/%s (%s %s) go/%s", cli.LibVersion, runtime.GOOS, runtime.GOARCH, runtime.Version())

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-Api-Key", cli.ApiKey)
	req.Header.Add("Accept", "application/json")
	for k, v := range libReq.headers {
		req.Header.Add(k, fmt.Sprint(v))
	}

	resp, err := cli.HttpClient.Do(req)
	if err != nil {
		//if Verbose {
		//	log
		//}
		return fmt.Errorf("makeRequest: error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		var nfError model.NotFoundError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&nfError); decodeErr != nil {
			return fmt.Errorf("status: %d - Error: error decoding the response for an HTTP error: %w", resp.StatusCode, decodeErr)
		}
		return nfError
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var restError model.RestError
		if decodeErr := json.NewDecoder(resp.Body).Decode(&restError); decodeErr != nil {
			return fmt.Errorf("status: %d - Error: error decoding the response for an HTTP error: %w", resp.StatusCode, decodeErr)
		}
		return restError
	}

	if out != nil {
		if err = json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("makeRequest: error while reading response body: %w", err)
		}
	}

	return nil
}

func requestLog(url string, libReq apiRequest) string {
	var log strings.Builder
	log.WriteString("URL: ")
	log.WriteString(fmt.Sprintf("'%s'", url))
	log.WriteRune('\n')
	log.WriteString("Method: ")
	log.WriteString(libReq.method)
	log.WriteRune('\n')
	log.WriteString("Body: ")
	log.WriteString(libReq.method)
	log.WriteRune('\n')
	log.WriteString("Headers: \n")
	for k, v := range libReq.headers {
		log.WriteString(k)
		log.WriteString(strings.Join(v, ","))
		log.WriteRune('\n')
	}

	return log.String()
}
