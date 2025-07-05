package handler

import (
	"github.com/gin-gonic/gin"
	"mail/model"
	"net/http"
)

func PostMail(c *gin.Context) {
	var newMail model.Mail

	if err := c.BindJSON(&newMail); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, newMail)
}
