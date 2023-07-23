package models

type Edition struct {
	Editioncode string `json:"editioncode"`
	Editionname string `json:"editionname"`
	Editiondate string `json:"editiondate"`
	Editiontype string `json:"editiontype"`
	MtgoCode    string `json:"mtgoCode"`
}
