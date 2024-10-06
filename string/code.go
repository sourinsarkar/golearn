package main
import ("fmt";"strings")

func main() {
	// var name1 string = "Sourin";
	// fmt.Println(name1)
	// fmt.Printf("%c\n",name1[5])
	// stringLength := len(name1)
	// fmt.Println(stringLength)

	// Compare(), Contains(), Replaces(), ToLower(), ToUpper(), Split()
	str1 := "Hello"
	str2 := "Hallo"
	str3 := "Hey"

	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare(str1, str3))
	fmt.Println(strings.Compare(str2, str3))

	str4 := "Mississippi"
	fmt.Println(strings.Replace(str4, "s", "t", 1))

	fmt.Println(strings.ToUpper("sourin"))
	fmt.Println(strings.ToLower("SOURIN"))

	str5 := "I love CS"
	fmt.Println(strings.Split(str5, " "))

	// Create String from slice
	stringSlice := []string { "I", "love", "India" }
	fmt.Println(strings.Join(stringSlice, " ") + ".")

	str6 := "Go"
	str6 = str6 + "lang"
	fmt.Println(str6)
}