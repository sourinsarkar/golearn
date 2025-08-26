package main
import "fmt"

func main (){
	var n int = 1

	for n <= 10 {
		fmt.Println(n)
		n++
	}

	n = 1

	for n <= 10 {
		var product int = 5 * n

		fmt.Printf("5 * %d = %d\n", n, product)
		n++
	}
}