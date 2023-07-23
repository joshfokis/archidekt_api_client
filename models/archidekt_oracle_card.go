package models

type OracleCard struct {
	ID             int            `json:"id"`
	Cmc            int            `json:"cmc"`
	ColorIdentity  []interface{}  `json:"colorIdentity"`
	Colors         []interface{}  `json:"colors"`
	Faces          []interface{}  `json:"faces"`
	Layout         string         `json:"layout"`
	ManaCost       string         `json:"manaCost"`
	ManaProduction ManaProduction `json:"manaProduction"`
	Name           string         `json:"name"`
	Power          string         `json:"power"`
	Salt           float64        `json:"salt"`
	SubTypes       []interface{}  `json:"subTypes"`
	SuperTypes     []interface{}  `json:"superTypes"`
	Text           string         `json:"text"`
	Tokens         []interface{}  `json:"tokens"`
	Toughness      string         `json:"toughness"`
	Types          []string       `json:"types"`
	Loyalty        interface{}    `json:"loyalty"`
}
