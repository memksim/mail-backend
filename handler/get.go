package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"mail/model"
	"net/http"
)

func GetMails(c *gin.Context, db *sql.DB) {
	cursor, err := db.Query("SELECT * FROM mails")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка сервера при получении данных",
		})

		log.Fatal("Ошибка запроса к БД:", err)
		return
	}

	defer func(cursor *sql.Rows) {
		if err := cursor.Close(); err != nil {
			log.Fatal("Ошибка закрытия курсора БД:", err)
		}
	}(cursor)

	var mails []model.Mail
	for cursor.Next() {
		var mail model.Mail
		if err := cursor.Scan(&mail); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Ошибка чтения данных",
			})

			log.Fatal("Ошибка чтения данных:", err)
			return
		}
		mails = append(mails, mail)
	}

	c.IndentedJSON(http.StatusOK, mails)
}
