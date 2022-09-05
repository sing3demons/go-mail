package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func sendGoMail(username string, password string, to []string) {
	//Get html
	var body bytes.Buffer
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(&body, struct{ Name string }{Name: "sing3demons"})

	// Send Gomail
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", to...)

	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// func sendMailSimple(subject string, body string, to []string) {
// 	username := os.Getenv("USER_GMAIL")
// 	password := os.Getenv("PASSWORD_GMAIL")
// 	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

// 	msg := "Subject: " + subject + "\n" + body

// 	err := smtp.SendMail(
// 		"smtp.gmail.com:587",
// 		auth,
// 		username,
// 		to,
// 		[]byte(msg),
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

// func sendMailSimpleHTML(subject string, name string, to []string) {
// 	username := os.Getenv("USER_GMAIL")
// 	password := os.Getenv("PASSWORD_GMAIL")

	// Get html
// 	var body bytes.Buffer
// 	t, err := template.ParseFiles("./index.html")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	t.Execute(&body, struct{ Name string }{Name: name})

// 	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

// 	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

// 	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

// 	err = smtp.SendMail(
// 		"smtp.gmail.com:587",
// 		auth,
// 		username,
// 		to,
// 		[]byte(msg),
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USER_GMAIL")
	password := os.Getenv("PASSWORD_GMAIL")
	to := os.Getenv("TO_MAIL")

	// sendMailSimple(
	// 	"Another subject",
	// 	"Another body",
	// 	[]string{to},
	// )
	// sendMailSimpleHTML("Another subject", "sing3demons", []string{to})
	sendGoMail(username, password, []string{to})
}
