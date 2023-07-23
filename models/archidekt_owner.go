package models

type Owner struct {
	ID          int           `json:"id"`
	Username    string        `json:"username"`
	Avatar      string        `json:"avatar"`
	Moderator   bool          `json:"moderator"`
	PledgeLevel interface{}   `json:"pledgeLevel"`
	Roles       []interface{} `json:"roles"`
}
