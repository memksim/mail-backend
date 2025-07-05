package handler

import (
	"github.com/gin-gonic/gin"
	"mail/model"
	"time"
)

func GetMails(c *gin.Context) {
	c.IndentedJSON(200, mails)
}

var mails = [3]model.Mail{
	{
		Id: 1,
		Sender: model.User{
			Id:        51,
			FirstName: "Ozon",
		},
		Title:      "Выгода сегодня!",
		Body:       "Купите товары по выгодной цене",
		Time:       time.Date(2025, time.July, 1, 12, 14, 00, 00, time.UTC).UnixMilli(),
		IsRead:     false,
		IsBookmark: false,
	},
	{
		Id: 2,
		Sender: model.User{
			Id:        17,
			FirstName: "Максим",
			LastName:  "Косенко",
		},
		Title:      "Оставил замечания",
		Body:       "Оставил замечания по пр, надо править",
		Time:       time.Date(2025, time.July, 4, 22, 17, 00, 00, time.UTC).UnixMilli(),
		IsRead:     false,
		IsBookmark: true,
	},
	{
		Id: 3,
		Sender: model.User{
			Id:        12,
			FirstName: "Катя",
			LastName:  "Косенко",
		},
		Title:      "Привет!",
		Body:       "Привет! Как дела?",
		Time:       time.Date(2025, time.July, 5, 10, 11, 00, 00, time.UTC).UnixMilli(),
		IsRead:     true,
		IsBookmark: false,
	},
}
