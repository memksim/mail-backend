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

	router.GET("/mails", func(context *gin.Context) {
		handler.GetMails(context, db)
	})

	router.POST("/mail", func(context *gin.Context) {
		handler.PostMail(context, db)
	})

	router.POST("/user", func(context *gin.Context) {
		handler.PostUser(context, db)
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
