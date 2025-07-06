package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"mail/model"
	"net/http"
	"strconv"
)

func ReadMail(context *gin.Context, db *sql.DB) {
	id, err := strconv.ParseInt(context.Query("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось получить численное представление id",
		})
		return
	}
	if _, err := db.Exec("UPDATE mails SET is_read = true WHERE id = ?", id); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка обновления данных",
		})
		log.Fatal("Ошибка обновления данных: ", err)
		return
	}

	query, err := db.Query("SELECT * FROM mails WHERE id = ?", id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения данных",
		})
		log.Fatal("Ошибка получения данных: ", err)
		return
	}
	var mails []model.Mail
	for query.Next() {
		var mail model.Mail
		if err := query.Scan(&mail.Id, &mail.Sender, &mail.Recipient, &mail.Title, &mail.Body, &mail.IsBookmark, &mail.IsRead, &mail.Time); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Ошибка получения данных",
			})
			log.Fatal("Ошибка получения данных: ", err)
		}
		log.Println("mail: ", mail)
		mails = append(mails, mail)
	}

	defer CloseCursor(query)

	if len(mails) > 0 {
		context.JSON(http.StatusOK, mails[0])
	}
}

func BookmarkMail(context *gin.Context, db *sql.DB) {
	isBookmark := context.Query("is_bookmark")
	if isBookmark != "true" && isBookmark != "false" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "is_bookmark должно быть 'true' или 'false'",
		})
		return
	}
	id, err := strconv.ParseInt(context.Query("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Не удалось получить численное представление id",
		})
		return
	}

	if _, err := db.Exec("UPDATE mails SET is_bookmark = ? WHERE id = ?", isBookmark, id); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка обновления данных",
		})
		log.Fatal("Ошибка обновления данных: ", err)
		return
	}

	query, err := db.Query("SELECT * FROM mails WHERE id = ?", id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка получения данных",
		})
		log.Fatal("Ошибка получения данных: ", err)
		return
	}
	var mails []model.Mail
	for query.Next() {
		var mail model.Mail
		if err := query.Scan(&mail.Id, &mail.Sender, &mail.Recipient, &mail.Title, &mail.Body, &mail.IsBookmark, &mail.IsRead, &mail.Time); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Ошибка получения данных",
			})
			log.Fatal("Ошибка получения данных: ", err)
		}
		mails = append(mails, mail)
	}

	defer CloseCursor(query)

	if len(mails) > 0 {
		context.JSON(http.StatusOK, mails[0])
	}
}
