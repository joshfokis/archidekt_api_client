package requests

import (
	"fmt"
	"io"
)

type GetDecksRequest struct {
	Page int
}

func NewGetDecksRequest(page int) *GetDecksRequest {
	return &GetDecksRequest{
		Page: page,
	}
}

func (req *GetDecksRequest) Method() string {
	return "GET"
}

func (req *GetDecksRequest) Path() string {
	return fmt.Sprintf("decks/?page=%d", req.Page)
}

func (req *GetDecksRequest) Body() (io.Reader, error) {
	return nil, nil
}
