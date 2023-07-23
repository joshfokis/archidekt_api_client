package models

type Categories struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	IncludedInDeck  bool   `json:"includedInDeck"`
	IncludedInPrice bool   `json:"includedInPrice"`
	IsPremier       bool   `json:"isPremier"`
}
