package common

import (
	"github.com/tekab-dev/tekab-test/models"
	"log"
	"net/smtp"
)

func Send(candidates []models.Candidate) {
	from := "omar.sastec@gmail.com"
	pass := "omarSASTEC123"
	for _, candidate := range candidates {
		to := candidate.Email
		name := candidate.Name

		msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" + "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Hello there\n\n" +
			"name" + name + "\n" +
			"<br><a href=\"https://www.w3schools.com\">here</a>"

		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))

		if err != nil {
			log.Printf("smtp error: %s", err)
			return
		}

		log.Print("sent, visit http://foobarbazz.mailinator.com")
	}
}
