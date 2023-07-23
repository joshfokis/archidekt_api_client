package models

// type ArchidektResponse struct {
// 	Count    int         `json:"count"`
// 	Next     string      `json:"next"`
// 	Previous interface{} `json:"previous"`
// 	Results  Results     `json:"results"`
// }

// type Owner struct {
// 	ID          int           `json:"id"`
// 	Username    string        `json:"username"`
// 	Avatar      string        `json:"avatar"`
// 	Moderator   bool          `json:"moderator"`
// 	PledgeLevel interface{}   `json:"pledgeLevel"`
// 	Roles       []interface{} `json:"roles"`
// }

// type Colors struct {
// 	W int `json:"W"`
// 	U int `json:"U"`
// 	B int `json:"B"`
// 	R int `json:"R"`
// 	G int `json:"G"`
// }

// type Results struct {
// 	ID             int           `json:"id"`
// 	Name           string        `json:"name"`
// 	CreatedAt      time.Time     `json:"createdAt"`
// 	UpdatedAt      time.Time     `json:"updatedAt"`
// 	DeckFormat     int           `json:"deckFormat"`
// 	Description    string        `json:"description"`
// 	Featured       string        `json:"featured"`
// 	CustomFeatured string        `json:"customFeatured"`
// 	Game           interface{}   `json:"game"`
// 	Private        bool          `json:"private"`
// 	ViewCount      int           `json:"viewCount"`
// 	Cards          []Card        `json:"cards"`
// 	Points         int           `json:"points"`
// 	UserInput      int           `json:"userInput"`
// 	Categories     Categories    `json:"categories"`
// 	CommentRoot    int           `json:"commentRoot"`
// 	Editors        []interface{} `json:"editors"`
// 	ParentFolder   int           `json:"parentFolder"`
// 	Bookmarked     bool          `json:"bookmarked"`
// 	DeckTags       []interface{} `json:"deckTags"`
// 	CardPackage    interface{}   `json:"cardPackage"`
// }

// type Categories struct {
// 	ID              int    `json:"id"`
// 	Name            string `json:"name"`
// 	IncludedInDeck  bool   `json:"includedInDeck"`
// 	IncludedInPrice bool   `json:"includedInPrice"`
// 	IsPremier       bool   `json:"isPremier"`
// }

// type ManaProduction struct {
// 	W interface{} `json:"W"`
// 	U interface{} `json:"U"`
// 	B interface{} `json:"B"`
// 	R interface{} `json:"R"`
// 	G interface{} `json:"G"`
// 	C interface{} `json:"C"`
// }

// type Card struct {
// 	ID              int           `json:"id"`
// 	Artist          string        `json:"artist"`
// 	TcgProductID    int           `json:"tcgProductId"`
// 	CkFoilID        int           `json:"ckFoilId"`
// 	CkNormalID      int           `json:"ckNormalId"`
// 	CmEd            string        `json:"cmEd"`
// 	CollectorNumber string        `json:"collectorNumber"`
// 	Multiverseid    int           `json:"multiverseid"`
// 	MtgoFoilID      int           `json:"mtgoFoilId"`
// 	MtgoNormalID    int           `json:"mtgoNormalId"`
// 	UID             string        `json:"uid"`
// 	Edition         Edition       `json:"edition"`
// 	Flavor          string        `json:"flavor"`
// 	Games           []interface{} `json:"games"`
// 	Options         []string      `json:"options"`
// 	OracleCard      OracleCard    `json:"oracleCard"`
// 	Owned           int           `json:"owned"`
// 	Rarity          string        `json:"rarity"`
// }

// type Edition struct {
// 	Editioncode string `json:"editioncode"`
// 	Editionname string `json:"editionname"`
// 	Editiondate string `json:"editiondate"`
// 	Editiontype string `json:"editiontype"`
// 	MtgoCode    string `json:"mtgoCode"`
// }

// type OracleCard struct {
// 	ID             int            `json:"id"`
// 	Cmc            int            `json:"cmc"`
// 	ColorIdentity  []interface{}  `json:"colorIdentity"`
// 	Colors         []interface{}  `json:"colors"`
// 	Faces          []interface{}  `json:"faces"`
// 	Layout         string         `json:"layout"`
// 	ManaCost       string         `json:"manaCost"`
// 	ManaProduction ManaProduction `json:"manaProduction"`
// 	Name           string         `json:"name"`
// 	Power          string         `json:"power"`
// 	Salt           float64        `json:"salt"`
// 	SubTypes       []interface{}  `json:"subTypes"`
// 	SuperTypes     []interface{}  `json:"superTypes"`
// 	Text           string         `json:"text"`
// 	Tokens         []interface{}  `json:"tokens"`
// 	Toughness      string         `json:"toughness"`
// 	Types          []string       `json:"types"`
// 	Loyalty        interface{}    `json:"loyalty"`
// }
