package lib

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gopkg.in/gomail.v2"
)

func SendMail(name string, email string, text string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	message := gomail.NewMessage()
	message.SetHeader("From", "Cots Dummy <cotsdummymail@gmail.com>")
	message.SetHeader("To", os.Getenv("EMAIL"))
	message.SetHeader("Subject", "CotsNepal Contact by "+name+" <"+email+">")
	message.SetBody("text/plain", text)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Email Sent Successfully.")
	}
}
