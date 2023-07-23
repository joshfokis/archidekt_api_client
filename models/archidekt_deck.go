package models

import "time"

type Deck struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	DeckFormat     int           `json:"deckFormat"`
	Description    string        `json:"description"`
	Featured       string        `json:"featured"`
	CustomFeatured string        `json:"customFeatured"`
	Game           interface{}   `json:"game"`
	Private        bool          `json:"private"`
	ViewCount      int           `json:"viewCount"`
	Cards          []Cards       `json:"cards"`
	Points         int           `json:"points"`
	UserInput      int           `json:"userInput"`
	Categories     []Categories  `json:"categories"`
	CommentRoot    int           `json:"commentRoot"`
	Editors        []interface{} `json:"editors"`
	ParentFolder   int           `json:"parentFolder"`
	Bookmarked     bool          `json:"bookmarked"`
	DeckTags       []interface{} `json:"deckTags"`
	CardPackage    interface{}   `json:"cardPackage"`
}
