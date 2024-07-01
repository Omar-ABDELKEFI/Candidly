package common

import (
	"github.com/tekab-dev/tekab-test/models"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

func Send(candidates []models.Candidate, testId uint) {
	from := "tekab.dev@gmail.com"
	pass := "tekab.dev123"
	for _, candidate := range candidates {
		to := candidate.Email
		log.Println(os.Getenv("key"), "os.Getenv(\"key\")os.Getenv(\"key\")")
		testId := strconv.FormatUint(uint64(testId), 10)
		candidateId := strconv.FormatUint(uint64(candidate.ID), 10)
		valueTestCandidateId := "testId=" + testId + "&&candidateId=" + candidateId
		TestCandidateIdEncrypted := AesEncrypt(valueTestCandidateId, os.Getenv("key"))
		idTestCandidate := AesDecrypt(TestCandidateIdEncrypted, os.Getenv("key"))
		log.Println(TestCandidateIdEncrypted, "TestCandidateIdEncrypted in common/mailer")
		log.Println(idTestCandidate, "idTestCandidate in common/mailer")
		link := "http://127.0.0.1:3000/quiz/" + TestCandidateIdEncrypted
		msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" + "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Test invitation\n\n" +
			"Hi,<br><br>We are pleased to invite you to our online test. Use the URL below to take the test from the comfort of your home   " +
			"<br><br><a href=" + link + ">test (click to start)</a><br><br> Good luck<br><br> Cordially, "

		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))

		if err != nil {
			log.Printf("smtp error: %s", err)
			return
		}

	}
}
