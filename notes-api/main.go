package main

import (
	"notes-api/handlers"
	"notes-api/middleware"
	"notes-api/models"
	"notes-api/utils"
	"notes-api/db"
	
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	r := gin.Default()

	r.POST("/register", Register)
	r.POST("/login", Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/notes", CreateNote)
		auth.GET("/my-notes", MyNotes)
		auth.GET("/notes/:id", GetNote)
		auth.PUT("/notes/:id", UpdateNote)
		auth.DELETE("/notes/:id", DeleteNote)
	}

	r.Run(":8080")
}