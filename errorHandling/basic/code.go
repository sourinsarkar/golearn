package main
import (
	"fmt"
	"errors"
)

// Using New() function
func checkNumber(n int) error {
	customError := errors.New("Invalid number")

	if n != 4 {
		return customError
	}

	return nil
}

func main() {
	number := 4

	if checkNumber(number) != nil {
		fmt.Println(checkNumber(number))
	} else {
		fmt.Println("Valid number")
	}
}

// we can also use Errorf() function which comes with fmt p ackage itself
// it is used similar to the way New() function used.