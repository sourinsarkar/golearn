package main
import "fmt"

func incrementNumber() func() int {
	n := 0

	return func() int {
		n++
		return n
	}
}

func main() {
	instance1 := incrementNumber()
	fmt.Println(instance1())
	fmt.Println(instance1())
	fmt.Println(instance1())
	fmt.Println(instance1())

	instance2 := incrementNumber()
	fmt.Println(instance2())
	fmt.Println(instance2())
}