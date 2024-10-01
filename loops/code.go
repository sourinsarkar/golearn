package main
import "fmt"

func main() {
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Println(i);
	}

	fmt.Println()
	
	numbers := [5] int { 11, 12, 13, 14, 15 }
	
	for item := range numbers {
		fmt.Printf("%d\n", numbers[item])
	}
	
	fmt.Println()
	
	for index, item := range numbers {
		fmt.Printf("%d %d\n", index, item)
	}
	
	fmt.Println()
	
	for _, item := range numbers {
		fmt.Printf("%d\n", item)
	}
	
	fmt.Println()
	
	for index := range numbers {
		fmt.Printf("%d\n", index)
	}
	
	fmt.Println()

	index := [5] int { 10, 20, 30, 40, 50 }

	for index := range index {
		fmt.Println(index)
	}
}