package endpoints

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	baseURL string
	apiKey  string
	token   string
}

func New(baseURL string, apiKey string) *Client {
	parsedURL := fmt.Sprintf("%s/auth/v1", baseURL)
	return &Client{
		client: http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: parsedURL,
		apiKey:  apiKey,
	}
}

func (c Client) WithCustomGoTrueURL(url string) *Client {
	return &Client{
		client:  c.client,
		baseURL: url,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

func (c Client) WithToken(token string) *Client {
	return &Client{
		client:  c.client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   token,
	}
}

func (c Client) WithClient(client http.Client) *Client {
	return &Client{
		client:  client,
		baseURL: c.baseURL,
		apiKey:  c.apiKey,
		token:   c.token,
	}
}

// Returns a copy of a HTTP client that will not follow redirects.
func noRedirClient(client http.Client) http.Client {
	return http.Client{
		Transport: client.Transport,
		Jar:       client.Jar,
		Timeout:   client.Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
