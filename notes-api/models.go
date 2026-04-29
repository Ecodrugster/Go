package main

type User struct {
	ID   int    `json:"id"`
	Email  string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	IsPublic bool `json:"is_public"`
	UserID int `json:"user_id"`
}
