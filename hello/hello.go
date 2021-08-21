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
	names := []string{"Sardor", "Khuzha", "Sardorkhuja"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
