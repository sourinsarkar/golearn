package main
import "fmt"

func main() {
	n := 1

	for n <= 10 {
		product := 5 * n
		
		fmt.Printf("5 * %d = %d\n", n, product)
		n++
	}
}	