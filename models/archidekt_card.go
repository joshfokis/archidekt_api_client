package models

type Cards struct {
	Card `json:"card"`
}

type Card struct {
	ID              int           `json:"id"`
	Artist          string        `json:"artist"`
	TcgProductID    int           `json:"tcgProductId"`
	CkFoilID        int           `json:"ckFoilId"`
	CkNormalID      int           `json:"ckNormalId"`
	CmEd            string        `json:"cmEd"`
	CollectorNumber string        `json:"collectorNumber"`
	Multiverseid    int           `json:"multiverseid"`
	MtgoFoilID      int           `json:"mtgoFoilId"`
	MtgoNormalID    int           `json:"mtgoNormalId"`
	UID             string        `json:"uid"`
	Edition         Edition       `json:"edition"`
	Flavor          string        `json:"flavor"`
	Games           []interface{} `json:"games"`
	Options         []string      `json:"options"`
	OracleCard      OracleCard    `json:"oracleCard"`
	Owned           int           `json:"owned"`
	Rarity          string        `json:"rarity"`
}
