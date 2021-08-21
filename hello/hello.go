package main

import (
	"fmt"
	"greetings"
	"log"
)

func init() {
	log.SetPrefix("error happened: ")
	log.SetFlags(0)
}

func main() {
	message, error := greetings.Hello("Sardor")
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
