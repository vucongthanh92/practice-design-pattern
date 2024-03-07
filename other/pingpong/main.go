package main

import (
	"fmt"
	"time"
)

func ping(ping chan string) {
	for {
		ping <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func pong(ping chan string, pong chan string) {
	for {
		select {
		case <-ping:
			pong <- "pong"
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {
	defer func() {
		fmt.Println("exit program...")
	}()

	var (
		pingChan = make(chan string, 1000)
		pongChan = make(chan string, 1000)
	)

	go ping(pingChan)
	go pong(pingChan, pongChan)

	for {
		select {
		case <-pingChan:
			fmt.Println(<-pingChan)
		case <-pongChan:
			fmt.Println(<-pongChan)
		}
	}
}
