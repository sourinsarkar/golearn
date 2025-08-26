package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetFlags(0)
	message, err := greetings.Hello("Sourin")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(message)
	}
}