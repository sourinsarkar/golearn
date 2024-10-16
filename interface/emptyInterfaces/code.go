package main

import (
	"fmt"
	"time"
)

func dikhao(i interface{}) {
	fmt.Println(i)
}

func dikhao2(kuchBhi string) {
	fmt.Println(kuchBhi)
}

func main() {
	vakya := "Namaste Sansar"
	sankhya := 69
	
	start := time.Now()
	dikhao(vakya)
	end := time.Now()
	
	fmt.Println("1st:", end.Sub(start))
	fmt.Println()
	
	start1 := time.Now()
	dikhao(sankhya)
	end1 := time.Now()
	
	fmt.Println("2nd:", end1.Sub(start1))
	fmt.Println()
	
	start2 := time.Now()
	dikhao2(vakya)
	end2 := time.Now()
	
	fmt.Println("3rd:", end2.Sub(start2))
}