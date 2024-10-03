package main
import "fmt"

func main() {
	friends := map[int]string{ 1: "Sourin", 2:"Swapnil" };
	fmt.Println("Initial map:", friends)

	// Element insertion
	friends[3] = "Suman"
	friends[4] = "Tamojeet"
	
	fmt.Println("Updated map:", friends)
	
	// Element deletion
	delete(friends, 4)

	fmt.Println("Updated map:", friends)

	// Map using make() function
	student := make(map[int]string)
	student[1] = "Potter"
	student[2] = "Wisley"
	student[3] = "Granger"

	fmt.Println(student)
}