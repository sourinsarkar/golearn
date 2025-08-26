package main
import "fmt"

func main() {

	// Explicit Type Casting

	var floatValue float32
	floatValue = 9.1

	var intValue int
	intValue = int(floatValue)

	fmt.Println(floatValue)
	fmt.Println(intValue)

	// Implicit Type Casting

	// There's no Implicit Type Casting
}