package main

import (
	"fmt"
	"net/http"

	entity "sammelanyogi/send-gomail.git/entity"
	lib "sammelanyogi/send-gomail.git/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", func(ctx *gin.Context) {
		var newMessage entity.Message
		if err := ctx.BindJSON(&newMessage); err != nil {
			fmt.Println(err)
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			lib.SendMail(newMessage.Name, newMessage.Email, newMessage.Text)
			ctx.IndentedJSON(http.StatusOK, newMessage)
		}

	})
	router.Run(":8080")
}
