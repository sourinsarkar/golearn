package main
import "fmt"

func main() {
	// Creating slize from am array
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	sliceNumbers := numbers[2 : 5]
	fmt.Print(sliceNumbers, "\n")

	
	newslice1 := []int{1,3,5,7}
	fmt.Print(newslice1, "\n")
	
	// Append on a slice
	newslice1 = append(newslice1, 9, 11)
	fmt.Print(newslice1, "\n")

	// Append a whole slice on to another
	newslice2 := []int{13, 15, 17, 19, 21}
	newslice1 = append(newslice1, newslice2...)
	fmt.Print(newslice1, "\n")

	// Copy one slice to another
	copyslice := []int{10, 20, 30}
	copy(copyslice, newslice2)
	fmt.Print(copyslice, "\n")
	fmt.Println("Length of copyslice:", len(copyslice))
	fmt.Println("Length of newslice2:", len(newslice2))

	// Create a Slice using make() function
	newslice3 := make([]int, 5, 7)
	newslice3[0] = 10
	newslice3[1] = 20
	newslice3[2] = 30
	newslice3[3] = 40
	newslice3[4] = 50
	// newslice3[5] = 60
	fmt.Println(newslice3)
	newslice3 = append(newslice3, 60, 70, 80)
	fmt.Println(newslice3)
	fmt.Println("Length:", len(newslice3))
	fmt.Println("Capacity:", cap(newslice3))
}