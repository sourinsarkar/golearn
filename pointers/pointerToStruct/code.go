package main
import "fmt"

func main() {
	type Student struct {
		name string
		age int
	}

	std1 := Student{ "Sourin", 12 }
	// var ptr *Student
	// ptr = &std1

	var ptr = &std1

	// fmt.Println(std1)
	fmt.Println(ptr)

	fmt.Println("Name:", ptr.name)
	fmt.Println("Age:", ptr.age)
	
	ptr.age++
	
	fmt.Println("Age:", ptr.age)
}