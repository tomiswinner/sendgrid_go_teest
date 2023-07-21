package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var FROM_EMAIL = os.Getenv("FROM_EMAIL")
var TO_EMAIL = os.Getenv("TO_EMAIL")

func main() {
	log.Print("function started")
	r := gin.Default()
	r.GET("/api/HttpTrigger1", func(c *gin.Context) {
		log.Print("HTTP Trigger Started")
		from := mail.NewEmail("Example User", FROM_EMAIL)
		subject := "Sending with SendGrid is Fun"
		to := mail.NewEmail("Example User", TO_EMAIL)
		plainTextContent := "and easy to do anywhere, even with Go"
		htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
		c.JSON(200, gin.H{
			"message2": "a",
		})
	})

	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	r.Run(":" + port)
}
