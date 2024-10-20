package main

import (
	"fmt"
	"sync"
)

func show(line string, wg *sync.WaitGroup) {
	defer wg.Done()

    fmt.Println(line)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    show("Hi there", &wg)
    go show("Hi Sourin", &wg)

    wg.Wait()
}