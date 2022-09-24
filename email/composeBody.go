package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func Compose() bytes.Buffer {
	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Diego Zanetti",
		Message: "Email de teste desenvolvido por Diego Zanetti.",
	})

	return body
}
