package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "api-nba-v1.p.rapidapi.com"
const endpoint = "https://api-nba-v1.p.rapidapi.com/"

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Key        string
	Logger     *log.Logger
}

func New(Key string, logger *log.Logger) (*Client, error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Key:        Key,
		Logger:     logger,
	}, nil
}

func (c *Client) GetRequestResult(ctx context.Context, method, relativePath string, querie string, respBody interface{}) (interface{}, error) {
	req, err := c.MakeRequest(ctx, http.MethodGet, relativePath, querie)
	if err != nil {
		return nil, err
	}
	code, err := c.DoRequest(req, &respBody)
	if err != nil {
		return 0, err
	}

	switch code {
	case http.StatusOK:
		return respBody, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	case http.StatusNotFound:
		return nil, fmt.Errorf("not found. user with id  may not exist")
	default:
		return nil, errors.New("unexpected error1")
	}
}

func (c *Client) MakeRequest(ctx context.Context, method, relativePath string, queries string) (*http.Request, error) {

	url := endpoint + relativePath
	if queries != "" {
		url = url + queries
	}
	req, _ := http.NewRequest("GET", url, nil)

	// set header
	req.Header.Add("x-rapidapi-host", baseURL)
	req.Header.Add("x-rapidapi-key", c.Key)

	return req, nil
}

func (c *Client) DoRequest(req *http.Request, respBody interface{}) (int, error) {
	resp, err := c.HTTPClient.Do(req)
	fmt.Println(resp)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return resp.StatusCode, nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(bodyBytes, respBody); err != nil {
		fmt.Printf("%#v\n", err)
		return 0, err
	}

	return resp.StatusCode, nil
}
