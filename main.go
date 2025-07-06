package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"mail/handler"
	"os"
)

var router *gin.Engine
var db *sql.DB

func main() {
	setupDb()
	setupRouter()
}

func setupRouter() {
	router = gin.Default()

	router.GET("/receivedMails", func(context *gin.Context) {
		handler.GetReceivedMails(context, db)
	})

	router.GET("/sentMails", func(context *gin.Context) {
		handler.GetSentMails(context, db)
	})

	router.POST("/mail", func(context *gin.Context) {
		handler.PostMail(context, db)
	})

	router.POST("/user", func(context *gin.Context) {
		handler.PostUser(context, db)
	})

	router.GET("/user", func(context *gin.Context) {
		handler.GetUserByEmail(context, db)
	})

	router.PATCH("/bookmarkMail", func(context *gin.Context) {
		handler.BookmarkMail(context, db)
	})

	router.PATCH("/readMail", func(context *gin.Context) {
		handler.ReadMail(context, db)
	})

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupDb() {
	var err error

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Ошибка взятия пути к проекту: ", err)
	}

	db, err = sql.Open("sqlite3", dir+"/mail")
	if err != nil {
		log.Fatal("Ошибка открытия соединения с базой данных: ", err)
	}
}
