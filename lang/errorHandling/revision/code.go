package main

import "fmt"

func main() {
	// defer ----------
	
	// defer fmt.Println("Hi five")
	// defer fmt.Println("Hi three")
	// defer fmt.Println("Hi four")
	// fmt.Println("Hi one")
	// fmt.Println("Hi two")
	
	// panic ----------

	// fmt.Println("Line1")
	// panic("Line2")
	// fmt.Println("Line3")

	msg := division(20,0)
	fmt.Println(msg)
}

func handlePanic() {
	detect := recover()

	if detect != nil {
		fmt.Println(detect)
	}
}

func division(number1, number2 int) int {
	defer handlePanic()
	
	if number2 == 0 {
		panic("Cannot divide a number by 0")
	} else {
		return (number1 / number2)
	}
}