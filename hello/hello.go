package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("error happened: ")
	log.SetFlags(0)

	message, error := greetings.Hello("")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
