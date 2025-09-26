package models

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type Post struct {
	ID int `json:"id"`
	userID int `json:"userID"`
	Title string `json:"title"`
	Content string `json:"content"`
}