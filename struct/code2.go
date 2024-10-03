package main
import "fmt"

// initializa the function Rectangle
type Rectangle func(int, int) int

// create a struct
type NewRectangle struct {
	length int
	breadth int
	rect Rectangle
}

func newRectExample() {
	rectInstance := NewRectangle {
		length: 10,
		breadth: 20,
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println(rectInstance.rect(rectInstance.length, rectInstance.breadth))
}
