package main

import (
	"fmt"
	"greetings"
)

func main() {
	message := greetings.Hello("Sardor")
	fmt.Println(message)
}
