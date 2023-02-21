package models

type Book struct{
	ID uint `json:"id" gorm:"primary_Key"`
	Title string `json:"title"`
	Author string `json:"author"`
}