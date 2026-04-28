package handlers

import (
	"2_TaskManager/db"
	"2_TaskManager/models"
	"2_TaskManager/utils"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

/* JSON input → hash password → store in DB */
func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Println(user)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	/* SQL command to INSERT new items on the table/db */
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"

	_, err = db.Conn.Exec(context.Background(), query, user.Username, hashedPassword)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		c.JSON(500, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered"})
}

/*
	FLOW
	login request

→ find user in DB
→ compare password
→ generate JWT
→ return token
*/
func Login(c *gin.Context) {

	var input models.User
	var dbUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}

	query := "SELECT id, username, password FROM users WHERE username=$1"

	err := db.Conn.QueryRow(context.Background(), query, input.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials!"})
		return
	}

	if !utils.CheckPassword(input.Password, dbUser.Password) {
		c.JSON(401, gin.H{"error": "Invalid Credentials!"})
		return
	}

	token, _ := utils.GenerateToken(dbUser.ID)

	c.JSON(200, gin.H{"token": token})
}
