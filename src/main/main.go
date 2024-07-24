package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Transaction struct {
	Amount   int    `json:"amount"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func main() {

	r := gin.Default()
	r.StaticFile("/", "../web/index.html")
	r.StaticFile("/send", "../web/send.html")
	r.StaticFile("/receive", "../web/receive.html")

	r.POST("/send", func(c *gin.Context) {
		var data Transaction
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		sendToken(data)
		c.JSON(http.StatusOK, gin.H{"message": "Tokens sent successfully"})
	})

	type portHttpData struct {
		Port int `json:"port"`
	}

	r.POST("/receive", func(c *gin.Context) {
		var data portHttpData
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		receiveToken(data.Port)

		c.JSON(http.StatusOK, gin.H{"message": "Received tokens successfully"})

	})

	r.Run(":8080")
}
