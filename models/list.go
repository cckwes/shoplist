package models

type List struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       int32  `json:"color"`
	UserID      string `json:"user_id"`
	Removed     bool   `json:"removed"`
	Done        bool   `json:"done"`
	Items       []Item `json:"items"`
}
