package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageBody struct {
	Message string `json:"message"`
}

type LinkBody struct {
	Link string `json:"link"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func Message(c *gin.Context, message string) {
	Success(c, MessageBody{Message: message})
}

func Link(c *gin.Context, link string) {
	Success(c, LinkBody{Link: link})
}
