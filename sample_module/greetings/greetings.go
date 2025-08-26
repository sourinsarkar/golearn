package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// message := fmt.Sprintf("Hi, %s", name)
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

func randomGreetings() string {
	messageTemplates := []string{
		"Hi, %s",
		"Hey, %s",
		"Hola, %s",
		"Hello, %s",
		"Namaste, %s",
		"Konnichiwa, %s",
	}

	return messageTemplates[rand.IntN(len(messageTemplates))]
}

func Hellos(names []string) (map[string]string, error) {
	newMessages := map[string]string{}

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		newMessages[name] = message
	}

	return newMessages, nil
}