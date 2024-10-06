package main
import "fmt"

func greet() {
	fmt.Println("Good Morning")
}

// Parameterised functions
func areaOfRectangle(length int, breadth int) {
	fmt.Println(length * breadth)
}

// with return
func areaOfSquare(side int) int {
	var area int = side * side
	// var areaInString string = string(area)
	return area
}

// Multi returns
func calculate(n1 int, n2 int) (int, int, int, int) {
	sum := n1 + n2
	difference := n1 - n2
	division := n1 / n2
	multiplication := n1 * n2

	return sum, difference, division, multiplication
}

func main() {
	greet()
	areaOfRectangle(4, 2)
	fmt.Println(areaOfSquare(4))
	
	sum, difference, division, multiplication := calculate(10, 10);
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Division:", division)
	fmt.Println("Multiplication:", multiplication)
}