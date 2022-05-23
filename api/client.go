package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultAPIVersion = "v5"
	defaultBaseURL    = "https://api.peopledatalabs.com/v5"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	apiVersion string
	//loggerLevel       // TODO
	//logger log.Logger // TODO: https://github.com/aws/aws-sdk-go/blob/107e8360115ce55fe47a88929a58eabe54849e29/aws/logger.go#L87
	libVersion string
}

func NewClient(apiKey, libVersion string, opts ...ClientOptions) Client {
	c := Client{
		apiKey:     apiKey,
		baseURL:    defaultBaseURL,
		apiVersion: defaultAPIVersion,
		libVersion: libVersion,
		httpClient: defaultHTTPClient(),
	}

	for _, opt := range opts {
		opt(&c)
	}

	return c
}

func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

type ClientOptions func(*Client)

// TODO: add docs
func WithHTTPClient(httpCli *http.Client) ClientOptions {
	return func(client *Client) {
		client.httpClient = httpCli
	}
}

// TODO: add docs
func WithTimeout(timeout time.Duration) ClientOptions {
	return func(client *Client) {
		client.httpClient.Timeout = timeout
	}
}

// TODO: WithLogLevel

func WithAPIVersion(version string) ClientOptions {
	return func(client *Client) {
		client.apiVersion = version
	}
}

type apiRequest struct {
	method  string
	path    string
	headers http.Header
	params  string
	body    io.Reader
}

func (cli Client) Get(ctx context.Context, path string, params interface{}, out interface{}) error {
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

func (cli Client) Post(ctx context.Context, path string, params interface{}, out interface{}) error {
	body, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("makeRequest: could not marshal JSON: %w", err)
	}

	req := apiRequest{
		method: http.MethodGet,
		path:   path,
		body:   bytes.NewBuffer(body),
	}
	req.headers.Set("Content-Type", "application/json")

	return cli.makeRequest(ctx, req, out)
}

func (cli Client) makeRequest(ctx context.Context, libReq apiRequest, out interface{}) error {
	urlStruct, err := url.Parse(cli.baseURL + libReq.path)
	if err != nil {
		return fmt.Errorf("makeRequest: %w", err)
	}
	urlStruct.RawQuery = libReq.params
	fmt.Println("URL:", urlStruct.String())

	// TODO: Verbose
	//	if Debug {
	//		fmt.Println(r.curlString(req, payload))
	//	}

	req, err := http.NewRequestWithContext(ctx, libReq.method, urlStruct.String(), libReq.body)
	if err != nil {
		return err
	}

	userAgent := fmt.Sprintf("peopledatalabs-go/%s (%s %s) go/%s", cli.libVersion, runtime.GOOS, runtime.GOARCH, runtime.Version())

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-Api-Key", cli.apiKey)
	req.Header.Add("Accept", "application/json")
	for k, v := range libReq.headers {
		req.Header.Add(k, fmt.Sprint(v))
	}

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		//if Verbose {
		//	log
		//}
		return fmt.Errorf("makeRequest: error making HTTP request: %w", err)
	}

	// TODO: Add custom error handling (errors.errorResponse)
	//switch {
	//case resp.StatusCode >= 400:
	//	//err = &HTTPError{}
	//	//if decodeErr := json.NewDecoder(resp.Body).Decode(err); decodeErr != nil {
	//	//	err = fmt.Errorf(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
	//	//	return nil, err
	//	//}
	//
	//	return err
	//}
	if resp.StatusCode != 200 {
		return fmt.Errorf("makeRequest: error making HTTP request: %w", err)
	}

	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("makeRequest: error while reading response body: %w", err)
	}

	return nil
}
