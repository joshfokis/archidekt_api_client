package api

import (
	"fmt"

	"archidekt/archidekt/models"
	"archidekt/archidekt/requests"
	"archidekt/archidekt/responses"
)

func ListAllDecks() ([]models.Deck, error) {
	resp, err := GetDecks(1)
	if err != nil {
		return nil, err
	}

	decks := resp.Decks
	pages := resp.Pages()

	for i := 2; i <= pages; i++ {
		resp, err := GetDecks(i)
		if err != nil {
			return nil, err
		}

		decks = append(decks, resp.Decks...)
		pages = resp.Pages()
	}
	return decks, nil
}

func ListUserDecks(username string) ([]models.Deck, error) {
	resp, err := GetUserDecks(username, 1)
	if err != nil {
		return nil, err
	}

	decks := resp.Decks
	pages := resp.Pages()

	for i := 2; i <= pages; i++ {
		fmt.Printf("Getting Page %d out of %d\r", i, pages)
		resp, err := GetUserDecks(username, i)
		if err != nil {
			return nil, err
		}

		decks = append(decks, resp.Decks...)
		pages = resp.Pages()
	}
	return decks, nil
}

func ListDeck(id int) (*responses.GetDeckResponse, error) {
	resp, err := GetDeck(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetUserDecks(username string, page int) (*responses.GetDecksResponse, error) {
	req := requests.NewGetUserDecksRequest(username, page)
	var result responses.GetDecksResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDecks(page int) (*responses.GetDecksResponse, error) {
	req := requests.NewGetDecksRequest(page)
	var result responses.GetDecksResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func GetDeck(id int) (*responses.GetDeckResponse, error) {
	req := requests.NewGetDeckRequest(id)
	var result responses.GetDeckResponse

	if err := GetClient().ExecuteRequest(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
