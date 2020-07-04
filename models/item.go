package models

type Item struct {
	Base
	Name     string `json:"name"`
	Category string `json:"category"`
	Count    uint   `json:"count"`
	Done     bool   `json:"done"`
	Removed  bool   `json:"removed"`
	ListID   string `json:"list_id"`
}
