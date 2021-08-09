package common

import (
	"github.com/tekab-dev/tekab-test/models"
	"log"
	"net/smtp"
	"strconv"
)

func Send(candidates []models.Candidate) {
	from := "tekab.dev@gmail.com"
	pass := "tekab.dev123"
	for _, candidate := range candidates {
		to := candidate.Email
		testId := strconv.FormatUint(uint64(candidate.Test[0].ID), 10)
		candidateId := strconv.FormatUint(uint64(candidate.ID), 10)
		link := "http://51.68.81.16:3000/testId=" + testId + "&&candidateId=" + candidateId
		msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" + "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Hello there\n\n" +
			"welcome to tekab test" +
			"<br><a href=" + link + ">here</a>"

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
