package middleware

import (
	"fmt"
	"net/smtp"
	)

func SendMail(subject string, message string) {
	from := "patrickjmccauley.dev@gmail.com"
	password := readCredFromFile("ecred.pickle")
	to := []string{
		"patrickjmccauley.dev@gmail.com",
	}

	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Build the body + subject
	msg := []byte("Subject:" + subject + "\n\n" + message)
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}
