package main
import "fmt"

func callByValue(num int) {
	// num = 30
	fmt.Println(num)
	fmt.Println(&num)
}

func callByRef(num *int) {
	// *num = 60
	fmt.Println(*num)
	fmt.Println(num)
}

func main() {
	foo := 2

	fmt.Printf("Value of foo: %d\n", foo) // Formats the values
	fmt.Printf("Value of foo: %v\n", foo) // Infers the value
	fmt.Printf("Value of foo: %d\n", &foo) // Formats the address as decimal integer
	fmt.Printf("Value of foo: %v\n", &foo) // Infers the address
	fmt.Printf("Value of foo: %p\n", &foo) // Original

	soo := 10
	fmt.Println(&soo)
	callByValue(soo)
	callByRef(&soo)
}