package main

import "fmt"

type Handler func() error

func execFunc(f Handler) error {
	return f()
}

type MessageHandler func(msg string) error

func HandlerAdapter(hdl MessageHandler, msg string) Handler {
	return func() error {
		return hdl(msg)
	}
}

func main() {
	f := func() error {
		fmt.Println("Hello World")
		return nil
	}

	execFunc(f)

	msgHdl := func(msg string) error {
		fmt.Println(msg)
		return nil
	}

	f = HandlerAdapter(msgHdl, "Hello World With Adapter Pattern")
	execFunc(f)
}
