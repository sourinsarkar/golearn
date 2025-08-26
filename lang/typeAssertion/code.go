package main
import "fmt"

func main() {
	var thing interface{}

	thing = 69
	fmt.Println(thing)
	// fmt.Println(thing.(string)) //just to check

	storeValue, boolValue := thing.(string)
	fmt.Println(storeValue)
	fmt.Println(boolValue)

	// Adding on to that
	if boolValue == true {
		fmt.Printf("It is a %t", !boolValue)
	} else {
		fmt.Printf("It is a %t", boolValue)
	}
}