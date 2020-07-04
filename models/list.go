package models

type List struct {
	Base
	Name    string `json:"name"`
	UserID  string `json:"user_id"`
	Removed bool   `json:"removed"`
	Done    bool   `json:"done"`
	Items   []Item `json:"items"`
}
