package models

type Book struct {
	ID     uint   `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
