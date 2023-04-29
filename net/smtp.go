package net

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/juandserrano/astro-cover/model"
)

func SendEmailNotification(n *model.Notification) {
  email := os.Getenv("GMAIL_ADDRESS")
  appPass := os.Getenv("GMAIL_ASTRO_COVER_APP_PASS")
  auth := smtp.PlainAuth("", email, appPass, "smtp.gmail.com")
	to := []string{email}

  var dataTable string
  for _, v := range n.Data {
    dataTable += fmt.Sprintf("%+v\n", v)
  }

	msg := []byte(fmt.Sprintf("To: %s\r\n" +
		"Subject: %s\r\n" +
    "\r\n" +
    "%s\n" +
		"%+v\r\n", email, n.Result, n.Day, dataTable))

	err := smtp.SendMail("smtp.gmail.com:587", auth, email, to, msg)
	if err != nil {
    log.Printf("smtp.SendMail error: %s", err)
	} else {
    log.Printf("Mail sent!")
  }
}
