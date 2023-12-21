package main

import (
	"fmt"
	"strings"
)

func main() {
	// create channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	fmt.Println("Type something and press ENTER (enter \"Q\" to quit)")
	for {
		// print prompt
		fmt.Println("-> ")

		// get user input
		var userInputString string
		_, _ = fmt.Scanln(&userInputString)

		if strings.ToLower(userInputString) == "q" {
			break
		}

		ping <- userInputString
		// wait for response
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("All done... Closing channels.")
	close(ping)
	close(pong)
}

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}
