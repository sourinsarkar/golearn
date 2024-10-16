package main
import "fmt"

func updateNumber(n *int) {
	*n = 30
}

func display() *int {
	n := 100
	return &n
}

func callByValue(n int) {
	n = 30
	fmt.Println(n)
}

func callByRef(n *int) {
	*n = 60
	fmt.Println("Address:",n)
	fmt.Println("Value:", *n)
}

func main() {

	// Pointer as function arguments
	number := 100

	fmt.Println("Number:", number)

	updateNumber(&number)
	fmt.Println("Number:", number)

	// Return pointers from function
	answer := display()
	fmt.Println("Address of Answer:", answer)
	fmt.Println("Value of Answer:", *answer)

	// Call by value and ref
	callByValue(20)
	temp := 40
	callByRef(&temp)
}