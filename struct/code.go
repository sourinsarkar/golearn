package main
import "fmt"

func main() {
	// declare a struct
	type Person struct {
		name string
		age int
	}

	// declare an instanc of the struct and initialize it
	p1 := Person{ "Sourin", 22 }
	fmt.Println(p1)

	// ----------------------------

	var p2 Person
	p2 = Person{ "Swapnil", 22 }
	fmt.Println(p2)

	// Another example
	type Rectangle struct {
		length int
		breadth int
	}

	var rect Rectangle
	rect = Rectangle{2,4}
	fmt.Println(rect)
	fmt.Println("Length:", rect.length)
	fmt.Println("Breadth:", rect.breadth)
	fmt.Println("Area:", rect.length * rect.breadth)
}