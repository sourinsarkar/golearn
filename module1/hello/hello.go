package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetFlags(0)
	// message, err := greetings.Hello("Sourin")

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(message)
	// }

	namesOfPeople := []string{"Milli", "Centi", "Deci", "Cule"}

	messages, err := greetings.Hellos(namesOfPeople)

	if err != nil {
		log.Fatal(err)
	} else {
		for _, message := range messages {
			fmt.Println(message)
		}
	}
}