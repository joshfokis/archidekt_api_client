package responses

import (
	"math"

	"archidekt/archidekt/models"
)

type GetDecksResponse struct {
	models.ArchidektResponse
	Decks []models.Deck `json:"results"`
}

func (resp GetDecksResponse) Pages() int {
	f := float64(resp.Count) / 20
	return int(math.Ceil(f))
}
