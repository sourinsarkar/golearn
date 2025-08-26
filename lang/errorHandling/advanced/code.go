package main

import "fmt"

func handlePanic() {
	rec := recover()

	if rec != nil {
		fmt.Println("RECOVER:", rec)
	}
}

func division(num1, num2 int) {
	defer handlePanic()

	if num2 == 0 {
		panic("Cannot divide a number by 0")
		// fmt.Println("Cannot divide a number by 0")
	} else {
		fmt.Println("Answer:", num1 / num2)
	}
}

func main() {
	// defer - Prevents execution of a function until the execution of other functions

	// defer fmt.Println("S")
	// defer fmt.Println("N")
	// defer fmt.Println("K")

	// panic - Panic statement to immediately end the execution of a program
	
	defer handlePanic()

	fmt.Println("Line 1")
	// fmt.Println("Line 2")
	
	division(8, 2)
	division(8, 0)
	division(4, 8)
	panic ("Its Panic!")

}