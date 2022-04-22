package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"zhouhg19@mails.tsinghua.edu.cn",
		"",
		"mails.tsinghua.edu.cn",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"mails.tsinghua.edu.cn:25",
		auth,
		"mails.tsinghua.edu.cn",
		[]string{"zhouhg19@mails.tsinghua.edu.cn"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
