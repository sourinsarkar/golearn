package main
import "fmt"

func main() {
	foo := 2

	fmt.Printf("Value of foo: %d\n", foo) // Formats the values
	fmt.Printf("Value of foo: %v\n", foo) // Infers the value
	fmt.Printf("Value of foo: %d\n", &foo) // Formats the address as decimal integer
	fmt.Printf("Value of foo: %v\n", &foo) // Infers the address
	fmt.Printf("Value of foo: %p\n", &foo) // Original
}