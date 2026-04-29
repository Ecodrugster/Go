package main

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := DB.QueryRow(
		"INSERT INTO users(email,password) VALUES($1,$2) RETURNING id",
		u.Email, u.Password,
	).Scan(&u.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var u User
	c.ShouldBindJSON(&u)

	var dbUser User
	err := DB.QueryRow(
		"SELECT id,password FROM users WHERE email=$1",
		u.Email,
	).Scan(&dbUser.ID, &dbUser.Password)

	if err != nil || dbUser.Password != u.Password {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token, _ := GenerateToken(dbUser.ID)
	c.JSON(200, gin.H{"token": token})
}

func CreateNote(c *gin.Context) {
	userID := c.GetInt("user_id")

	var n Note
	c.ShouldBindJSON(&n)

	err := DB.QueryRow(
		"INSERT INTO notes(title,content,is_public,user_id) VALUES($1,$2,$3,$4) RETURNING id",
		n.Title, n.Content, n.IsPublic, userID,
	).Scan(&n.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

func MyNotes(c *gin.Context) {
	userID := c.GetInt("user_id")

	rows, _ := DB.Query("SELECT id,title,content,is_public FROM notes WHERE user_id=$1", userID)

	var notes []Note
	for rows.Next() {
		var n Note
		rows.Scan(&n.ID, &n.Title, &n.Content, &n.IsPublic)
		notes = append(notes, n)
	}

	c.JSON(200, notes)
}

func GetNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")

	var n Note
	err := DB.QueryRow(
		"SELECT id,title,content,is_public,user_id FROM notes WHERE id=$1",
		id,
	).Scan(&n.ID, &n.Title, &n.Content, &n.IsPublic, &n.UserID)

	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	if n.UserID != userID && !n.IsPublic {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	c.JSON(200, n)
}

func UpdateNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")

	var ownerID int
	DB.QueryRow("SELECT user_id FROM notes WHERE id=$1", id).Scan(&ownerID)

	if ownerID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	var n Note
	c.ShouldBindJSON(&n)

	DB.Exec(
		"UPDATE notes SET title=$1, content=$2, is_public=$3 WHERE id=$4",
		n.Title, n.Content, n.IsPublic, id,
	)

	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteNote(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetInt("user_id")

	var ownerID int
	DB.QueryRow("SELECT user_id FROM notes WHERE id=$1", id).Scan(&ownerID)

	if ownerID != userID {
		c.JSON(403, gin.H{"error": "forbidden"})
		return
	}

	DB.Exec("DELETE FROM notes WHERE id=$1", id)
	c.JSON(200, gin.H{"message": "deleted"})
}