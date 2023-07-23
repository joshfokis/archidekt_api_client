package requests

import (
	"fmt"
	"io"
)

type GetDeckRequest struct {
	Id int
}

func NewGetDeckRequest(id int) *GetDeckRequest {
	return &GetDeckRequest{Id: id}
}

func (req *GetDeckRequest) Method() string {
	return "GET"
}

func (req *GetDeckRequest) Path() string {
	return fmt.Sprintf("decks/%d/", req.Id)
}

func (req *GetDeckRequest) Body() (io.Reader, error) {
	return nil, nil
}
