package greetings

import "fmt"

func Hello(name string) string {
	var message string = fmt.Sprintf("Hello, %s", name)
	return message
}
