package main

import (
	"fmt"
	"time"
)

func ping(input chan string, output chan string) {
	for {
		time.Sleep(1 * time.Second)
		message := <-input
		if message == "pong" {
			fmt.Println(message)
			output <- "ping"
		}
	}
}

func pong(input chan string, output chan string) {

	for {
		message := <-input
		if message == "ping" {
			fmt.Println(message)
			output <- "pong"
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {

	pingChan := make(chan string)
	pongChan := make(chan string)

	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)
	// kick it off
	pingChan <- "pong"

	for {

	}
}
