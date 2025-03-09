package models

type Category struct {
	Id   int64  `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}
