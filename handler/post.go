package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"mail/model"
	"net/http"
)

func PostMail(c *gin.Context, db *sql.DB) {
	var newMail model.Mail

	if err := c.BindJSON(&newMail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка чтения аргументов",
		})
		return
	}

	var recipient = newMail.Recipient

	//TODO: Also check if sender is in current session

	cursor, err := db.Query("SELECT * FROM users WHERE email = ?", recipient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка чтения получателя",
		})

		log.Fatal("Ошибка чтения получателя:", err)
		return
	}

	if !cursor.Next() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	} else {
		closeCursor(cursor)
	}

	defer closeCursor(cursor)

	result, err := db.Exec(
		"INSERT INTO mails (id, sender_email, recipient_email, title, body, is_bookmark, is_read, time) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		newMail.Id, newMail.Sender, newMail.Recipient, newMail.Title, newMail.Body, newMail.IsBookmark, newMail.IsRead, newMail.Time,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка записи данных",
		})

		log.Fatal("Ошибка записи данных:", err)
		return
	}

	if _, err := result.LastInsertId(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения записанного значения",
		})

		log.Fatal("Ошибка получения записанного значения:", err)
		return
	}

	c.JSON(http.StatusOK, newMail)
}

func PostUser(c *gin.Context, db *sql.DB) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка чтения аргументов",
		})
		return
	}

	result, err := db.Exec(
		"insert into users(email, first_name, last_name, avatar_url) values (?, ?, ?, ?)",
		newUser.Email, newUser.FirstName, newUser.LastName, newUser.Avatar,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка записи данных",
		})

		log.Fatal("Ошибка записи данных:", err)
		return
	}

	if _, err := result.LastInsertId(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения записанного значения",
		})

		log.Fatal("Ошибка получения записанного значения:", err)
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func closeCursor(cursor *sql.Rows) {
	if err := cursor.Close(); err != nil {
		log.Fatal("Ошибка закрытия курсора БД:", err)
	}
}
