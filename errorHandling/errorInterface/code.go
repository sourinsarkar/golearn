package main
import "fmt"

/*
	type error interface {
		Error() string
	}
*/

type DivisionByZero struct {
	message string
}

func (d DivisionByZero) Error() string {
	return "There is an error!"
}

func divide(num1 int, num2 int) (int, error) { 
	if num2 == 0 {
		return num2, &DivisionByZero{}
	} else {
		return num1 / num2, nil
	}
}

func main() {
	num1 := 20
	num2 := 0

	result, err := divide(num1, num2)

	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}