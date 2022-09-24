package main

import (
	"crypto/tls"
	"errors"
	"net"
	"net/smtp"
	"os"
)

type ServerConfig struct {
	from     string
	smtpHost string
	smtpPort string
	smtpAuth smtp.Auth
}

type loginAuth struct {
	username string
	password string
}

func ConfigureConn() ServerConfig {
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	conn, err := net.Dial("tcp", smtpHost+":"+smtpPort)
	if err != nil {
		println(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		println(err)
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		println(err)
	}

	auth := LoginAuth(from, password)

	if err = c.Auth(auth); err != nil {
		println(err)
	}

	server := ServerConfig{}

	server.from = from
	server.smtpHost = smtpHost
	server.smtpPort = smtpPort
	server.smtpAuth = auth

	return server
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
