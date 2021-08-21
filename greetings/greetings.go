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
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
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
