package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	var message string = fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting() string {
	formats := []string{
		"Привет, %s!",
		"Guten Tag, %s!",
		"Hello %s!",
	}

	return formats[rand.Intn(len(formats))]
}
