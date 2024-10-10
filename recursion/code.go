package main
import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}
	return n + sum(n - 1)
}

func facotrial(n int) int {
	if(n == 0) {
		return 1;
	}

	return n * facotrial(n - 1);
}

func displayNumber() func() int {
	number := 10

	return func() int {
		number++
		return number
	}
}

func main() {
	n := 0
	// fmt.Println("Sum:", sum(n))
	fmt.Println("Factorial:", facotrial(n))
	
	dn := displayNumber()
	fmt.Println(dn())
}