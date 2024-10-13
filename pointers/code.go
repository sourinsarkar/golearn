package main
import "fmt"

func main() {
	// n := 5
	// fmt.Printf("Value: %d\n", n)
	// fmt.Printf("Address of n: %p\n", &n)
	
	// var ptr *int = &n
	// fmt.Printf("Address of n: %p\n", ptr)

	// fmt.Println("Value of n:", ptr)
	
	// var num  int
	// var ptr *int
		
	// -------------------------------------------------------------
	
	// num = 22
	// fmt.Println("Address of num:",&num)
	// fmt.Println("Value of num:",num)
	
	// ptr = &num
	// fmt.Println("\nAddress of pointer ptr:",ptr)
	// fmt.Println("Content of pointer ptr:",*ptr)
	
	// num = 11
	// fmt.Println("\nAddress of pointer ptr:",ptr)
	// fmt.Println("Content of pointer ptr:",*ptr)
		
	// *ptr = 2
	// fmt.Println("\nAddress of num:",&num)
	// fmt.Println("Value of num:",num)

	// -------------------------------------------------------------
	
	var ptr1 *int
	fmt.Println(ptr1) // Output: <int>
	
	// -------------------------------------------------------------
	
	// A different way to initialize pointers in Go
	
	var ptr2 = new(int)
	fmt.Println(ptr2) // Output: 0xc00008c0d0
	
	// -------------------------------------------------------------

	// Pointer creation without * operator

	var something string = "Sourin"
	var pointer = &something

	fmt.Println(&something)
	fmt.Println(pointer)
}