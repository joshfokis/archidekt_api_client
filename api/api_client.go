package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"archidekt/archidekt/requests"
)

type ApiClient struct {
	Http *http.Client
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		Http: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   60 * time.Second,
		},
	}
}

var client *ApiClient = NewApiClient()

func GetClient() *ApiClient {
	return client
}

func (c *ApiClient) ExecuteRequest(req requests.Request, resp any) error {
	r, err := requests.Execute(req, c.Http)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, resp); err != nil {
		return err
	}

	return nil
}
