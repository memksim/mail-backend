package main

import (
	"github.com/gin-gonic/gin"
	"mail/handler"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.GET("/mails", handler.GetMails)

	router.POST("/create", handler.PostMail)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
