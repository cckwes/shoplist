package models

type Item struct {
	Base
	Name     string
	Category string
	Count    uint
	UserID   string
}
