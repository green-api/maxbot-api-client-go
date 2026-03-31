package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	e "github.com/green-api/maxbot-api-client-go/pkg/maxerr"
	"github.com/green-api/maxbot-api-client-go/pkg/models"

	"golang.org/x/time/rate"
)

type Config struct {
	BaseURL   string
	Token     string
	Timeout   time.Duration
	GlobalRPS int
}

type Client struct {
	baseURL       string
	token         string
	httpClient    *http.Client
	timeout       time.Duration
	globalLimiter *rate.Limiter
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.GlobalRPS <= 0 {
		cfg.GlobalRPS = 25 /* Exceeding the limit will lead to account ban */
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 35 * time.Second
	}

	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("BaseURL is not set")
	}

	if cfg.Token == "" {
		return nil, fmt.Errorf("Token is not set")
	}

	client := &Client{
		baseURL:       cfg.BaseURL,
		token:         cfg.Token,
		httpClient:    &http.Client{Timeout: cfg.Timeout},
		globalLimiter: rate.NewLimiter(rate.Limit(cfg.GlobalRPS), cfg.GlobalRPS),
	}

	return client, nil
}

type requestOptions struct {
	Query       url.Values
	Payload     any
	ContentType string
}

type RequestOption func(*requestOptions)

func WithQuery(params any) RequestOption {
	return func(r *requestOptions) {
		if v, ok := params.(url.Values); ok {
			r.Query = v
		} else {
			r.Query = models.ToValues(params)
		}
	}
}

func WithBody(body any) RequestOption {
	return func(o *requestOptions) {
		o.Payload = body
	}
}

func WithMultipart(fileName string, fileReader io.Reader) RequestOption {
	return func(r *requestOptions) {
		var body bytes.Buffer
		writer := multipart.NewWriter(&body)

		part, _ := writer.CreateFormFile("file", fileName)
		io.Copy(part, fileReader)
		writer.Close()

		r.Payload = &body
		r.ContentType = writer.FormDataContentType()
	}
}

func (c *Client) Request(ctx context.Context, method, path string, opts ...RequestOption) ([]byte, error) {
	options := &requestOptions{}
	for _, opt := range opts {
		opt(options)
	}

	var u *url.URL
	var err error

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		u, err = url.Parse(path)
	} else {
		base, parseErr := url.Parse(c.baseURL)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid base URL: %w", parseErr)
		}
		u = base.JoinPath(path)
	}

	if err != nil {
		return nil, fmt.Errorf("parse URL error: %w", err)
	}

	if len(options.Query) > 0 {
		u.RawQuery = options.Query.Encode()
	}

	var body io.Reader
	if options.Payload != nil {
		if r, ok := options.Payload.(io.Reader); ok {
			body = r
		} else {
			b, err := json.Marshal(options.Payload)
			if err != nil {
				return nil, fmt.Errorf("marshal error: %w", err)
			}
			body = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("create request error: %w", err)
	}

	if options.ContentType != "" {
		req.Header.Set("Content-Type", options.ContentType)
	}

	return c.do(req)
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	if err := c.globalLimiter.Wait(req.Context()); err != nil {
		return nil, fmt.Errorf("Rate limiter error: %w", err)
	}

	req.Header.Set("Authorization", c.token)
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := e.HandleErrorResponse(resp, body)
		return nil, fmt.Errorf("[%s %s] %w", req.Method, req.URL.Path, apiErr)
	}

	return body, nil
}

func (c *Client) SetGlobalRateLimit(rps int) {
	c.globalLimiter.SetLimit(rate.Limit(rps))
	c.globalLimiter.SetBurst(rps)
}

func (c *Client) GetTimeout() int {
	return int(c.timeout)
}
