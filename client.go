package addyrest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

const addyBaseURL = "https://app.addy.io"

type Client struct {
	BaseURL string
	Token   string
	Ctx     context.Context
}

// Initialize a new API client.
func NewCustomClient(ctx context.Context, url, token string) *Client {
	return &Client{
		BaseURL: url,
		Token:   token,
		Ctx:     ctx,
	}
}

// Initialize a new default API client, only provide the token.
func NewClient(token string) *Client {
	return NewCustomClient(context.Background(), addyBaseURL, token)
}

func getWithParams[T any](client *Client, apiPath string, params []string) (*T, error) {
	return get[T](client, apiPath+"?"+strings.Join(params, "&"))
}

func executeRequest(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, errors.New(res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func processContext(ctx context.Context) (context.Context, context.CancelFunc) {
	timeout := 30 * time.Second
	return context.WithTimeout(ctx, timeout)
}

func get[T any](client *Client, apiPath string) (*T, error) {
	reqContext, cancel := processContext(client.Ctx)
	defer cancel()

	url := client.BaseURL + "/" + apiPath
	req, err := http.NewRequestWithContext(reqContext, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	return parseJSON[T](body)
}

func delete[T any](client *Client, apiPath string) (*T, error) {
	reqContext, cancel := processContext(client.Ctx)
	defer cancel()

	url := client.BaseURL + "/" + apiPath
	req, err := http.NewRequestWithContext(reqContext, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, err
	}
	return parseJSON[T](body)
}

func post[T any](client *Client, apiPath string, data any) (*T, error) {
	timeout := 30 * time.Second
	reqContext, cancel := context.WithTimeout(client.Ctx, timeout)
	defer cancel()

	json, _ := toJSON(data)
	byteReader := bytes.NewReader(json)
	url := client.BaseURL + "/" + apiPath
	req, err := http.NewRequestWithContext(reqContext, "POST", url, byteReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	return parseJSON[T](body)
}

func patch[T any](client *Client, apiPath string, data any) (*T, error) {
	timeout := 30 * time.Second
	reqContext, cancel := context.WithTimeout(client.Ctx, timeout)
	defer cancel()

	json, _ := toJSON(data)
	byteReader := bytes.NewReader(json)
	url := client.BaseURL + "/" + apiPath
	req, err := http.NewRequestWithContext(reqContext, "PATCH", url, byteReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+client.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	return parseJSON[T](body)
}

func parseJSON[T any](s []byte) (*T, error) {
	var r T

	err := json.Unmarshal(s, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func toJSON(T any) ([]byte, error) {
	return json.Marshal(T)
}
