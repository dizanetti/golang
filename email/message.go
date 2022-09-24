package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
)

type Message struct {
	To          []string
	CC          []string
	BCC         []string
	Body        bytes.Buffer
	Attachments map[string][]byte
}

func NewMessage(body bytes.Buffer) *Message {
	return &Message{Body: body, Attachments: make(map[string][]byte)}
}

func (m *Message) AttachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) GetTo() []string {
	return m.To
}

func (m *Message) SetTo(to []string) {
	m.To = to
}

func (m *Message) GetCc() []string {
	return m.CC
}

func (m *Message) SetCC(cc []string) {
	m.CC = cc
}

func (m *Message) GetBcc() []string {
	return m.BCC
}

func (m *Message) SetBcc(bcc []string) {
	m.BCC = bcc
}

func (m *Message) GetBody() bytes.Buffer {
	return m.Body
}

func (m *Message) GetBodyByte() []byte {
	return m.Body.Bytes()
}

func (m *Message) SetBody(body bytes.Buffer) {
	m.Body = body
}

func (m *Message) GetAttachments() map[string][]byte {
	return m.Attachments
}

func (m *Message) SetAttachments(attachments map[string][]byte) {
	m.Attachments = attachments
}
