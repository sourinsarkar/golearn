package main
import "fmt"

func main() {
	var string1 string
	var int1 int

	fmt.Print("Enter STRING1 and INT1:")
	fmt.Scan(&string1, &int1)

	fmt.Printf("STRING1 %s and INT1 %d",string1,int1)
}