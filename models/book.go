package models

type Book struct {
    ID          uint   `json:"id" gorm:"primary_key"`
    Title       string `json:"title"`
    Author      string `json:"author"`
    Publisher   string `json:"publisher"`
    PublishYear int    `json:"publish_year"`
}