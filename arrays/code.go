package main
import "fmt"

func main() {
	// Array of defined size
	var definedArray = [5] int {10, 20, 30, 40, 50}
	fmt.Println(definedArray)

	// Array of undefined size
	var undefinedArray = [...] int {100, 200, 300, 400, 500}
	fmt.Println(undefinedArray)

	// Shorthand to declare an Array
	shorthandArray := [...] string {"Hi", "there"}
	fmt.Println(shorthandArray)

	// Accessing the elements of an array
	fmt.Println(shorthandArray[0])
	fmt.Println(shorthandArray[1])
	fmt.Printf("%s %s, how are you?\n", shorthandArray[0], shorthandArray[1])

	// Initialize the elements of 0 and 3 only
	newArray := [...]int{0:199, 3:499}
	fmt.Println(newArray)
	
	// Initialize the elements of 0 and 3 only in strings
	newStringArray := [...]string{1: "Sourin", 20:"Sarkar"}
	fmt.Println(newStringArray)

	// Length of an array in Go
	fmt.Println(len(undefinedArray))
	fmt.Println(len(newStringArray))

	// Looping through an array
	i := 0

	for i < len(undefinedArray) {
		fmt.Println(undefinedArray[i])
		i++
	}

	// Two Dimensional Array
	// array1 := [4]int{10, 20, 30, 40}
	// array2 := [4]int{100, 200, 300, 400}
	array3 := [2][2]int{{ 10, 20 }, {30, 40}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(array3[i][j], " ")
		}
		fmt.Println()
	}
	
}