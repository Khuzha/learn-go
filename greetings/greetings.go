package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	var message string = fmt.Sprintf("Hello, %s", name)
	return message, nil
}
