package main

import (
	"fmt"
	"greetings"
)

func main() {
	var message string = greetings.Hello("Sardor")
	fmt.Println(message)
}
