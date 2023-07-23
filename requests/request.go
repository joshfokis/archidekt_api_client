package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"archidekt/archidekt/responses"
)

type Request interface {
	Method() string
	Path() string
	Body() (io.Reader, error)
}

type RequestDoer interface {
	Do(*http.Request) (*http.Response, error)
}

func Execute(r Request, d RequestDoer) (*http.Response, error) {
	req := requestObject(r)

	if req == nil {
		return nil, fmt.Errorf("Could Not Build Request %v", r)
	}

	resp, err := d.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return handleHttpError(resp, r, d)
	}

	return resp, nil
}

func requestObject(r Request) *http.Request {
	url := urlFromPath(r.Path())

	body, err := r.Body()
	if err != nil {
		return nil
	}

	req, err := http.NewRequest(r.Method(), url, body)
	if err != nil {
		return nil
	}

	return req
}

func urlFromPath(path string) string {
	return fmt.Sprintf("https://archidekt.com/api/%s", path)
}

func handleHttpError(resp *http.Response, r Request, d RequestDoer) (*http.Response, error) {
	switch resp.StatusCode {
	case 429:
		return retryAfter(resp, r, d)
	case 502:
		return recoverFromDdos(r, d)
	default:
		defer resp.Body.Close()
		return resp, NewGenericHttpError(resp)
	}
}

func retryAfter(resp *http.Response, r Request, d RequestDoer) (*http.Response, error) {
	o, err := responses.NewErrorResponse(resp)
	if err != nil {
		return nil, err
	}

	timeout := o.Error.Data["retryAfter"].(int)
	timeToWait := time.Duration(timeout) * time.Second
	time.Sleep(timeToWait)

	return Execute(r, d)
}

func recoverFromDdos(r Request, d RequestDoer) (*http.Response, error) {
	timeToWait := time.Duration(10) * time.Second
	time.Sleep(timeToWait)

	return Execute(r, d)
}

func marshal(r Request) (io.Reader, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(body), nil
}
