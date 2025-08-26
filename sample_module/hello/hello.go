package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("Sourin")

	if err != nil {
		log.SetFlags(0)
		log.Fatal(err)
	} else {	
		fmt.Println(message)
	}
}