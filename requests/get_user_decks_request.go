package requests

import (
	"fmt"
	"io"
)

type GetUserDecksRequest struct {
	Username string
	Page     int
}

func NewGetUserDecksRequest(username string, page int) *GetUserDecksRequest {
	return &GetUserDecksRequest{
		Username: username,
		Page:     page,
	}
}

func (req *GetUserDecksRequest) Method() string {
	return "GET"
}

func (req *GetUserDecksRequest) Path() string {
	return fmt.Sprintf("decks/?owner=%s&exactowner=true&page=%d", req.Username, req.Page)
}

func (req *GetUserDecksRequest) Body() (io.Reader, error) {
	return nil, nil
}
