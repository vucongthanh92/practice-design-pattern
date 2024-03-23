package main

import "fmt"

type ChatMessage struct {
	Content string
	Sender  *Sender
}

type Sender struct {
	Name   string
	Avatar []byte
}

type SenderFactory struct {
	cacheSender map[string]*Sender
}

func (sf *SenderFactory) GetSender(name string) *Sender {
	return sf.cacheSender[name]
}

func main() {
	sender := SenderFactory{cacheSender: map[string]*Sender{
		"Peter": {
			Name:   "Peter Pan",
			Avatar: make([]byte, 1024*300),
		},
		"Mary": {
			Name:   "Mary Jame",
			Avatar: make([]byte, 1024*400),
		},
	}}

	fmt.Println([]ChatMessage{
		{
			Content: "reply 1",
			Sender:  sender.GetSender("Peter"),
		},
		{
			Content: "reply 2",
			Sender:  sender.GetSender("Mary"),
		},
		{
			Content: "reply 3",
			Sender:  sender.GetSender("Peter"),
		},
		{
			Content: "reply 4",
			Sender:  sender.GetSender("Mary"),
		},
	})
}
