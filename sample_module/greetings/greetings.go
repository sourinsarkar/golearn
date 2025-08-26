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
		"Hello, %s",
		"Hola, %s",
		"Namaste, %s",
		"Konnichiwa, %s",
	}

	return messageTemplates[rand.IntN(len(messageTemplates))]
}