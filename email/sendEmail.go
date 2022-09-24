package main

import (
	"fmt"
	"net/smtp"
)

func Send(message Message) {
	serverConfig := ConfigureConn()

	err := smtp.SendMail(
		serverConfig.smtpHost+":"+serverConfig.smtpPort,
		serverConfig.smtpAuth,
		serverConfig.from,
		message.To,
		message.GetBodyByte())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent!")
}
