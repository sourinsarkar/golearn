package main
import "fmt"

func main() {
	// create a map
	subjectMarks := map[string] float32 {"Sourin": 10, "Sky": 20, "Stark": 30}

	for subject, mark := range subjectMarks {
		fmt.Println(subject, ":", mark)
	}
}