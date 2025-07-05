package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"mail/config"
	"mail/model"
	"net/http"
)

func GetReceivedMails(c *gin.Context, db *sql.DB) {
	cursor, err := db.Query("SELECT * FROM mails where recipient_email = ?", config.CurrentUser.Email)
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

	var mails = make([]model.Mail, 10)
	for cursor.Next() {
		var mail model.Mail

		if err := cursor.Scan(&mail.Id, &mail.Sender, &mail.Recipient, &mail.Title, &mail.Body, &mail.IsBookmark, &mail.IsRead, &mail.Time); err != nil {
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

func GetSentMails(c *gin.Context, db *sql.DB) {
	cursor, err := db.Query("SELECT * FROM mails where sender_email = ?", config.CurrentUser.Email)
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

		if err := cursor.Scan(&mail.Id, &mail.Sender, &mail.Recipient, &mail.Title, &mail.Body, &mail.IsBookmark, &mail.IsRead, &mail.Time); err != nil {
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

func GetUserByEmail(c *gin.Context, db *sql.DB) {
	email := c.Query("email")

	cursor, err := db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось получить данные о пользователях",
		})

		log.Fatal("Не удалось получить данные о пользователях: ", err)
	}

	var users []model.User
	for cursor.Next() {
		var user model.User

		if err := cursor.Scan(&user.Email, &user.FirstName, &user.LastName, &user.Avatar); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Не удалось получить данные о пользователе",
			})
			log.Fatal("Не удалось получить данные о пользователе: ", err)
		}

		users = append(users, user)
	}

	defer CloseCursor(cursor)

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Пользователь не найден",
		})
	} else {
		c.IndentedJSON(http.StatusOK, users[len(users)-1])
	}
}
