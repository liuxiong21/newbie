package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 1000
	product := 0
	swapAndProduct(&a, &b, &product)
	fmt.Println(a, b, product)
	createArray()
	createSlice()
	fileName := "/opt/book/goeg/src/guess_separator/information.log"
	separatorName := GuessSeparator(fileName)
	fmt.Printf("Found separator %v in file %s\n", separatorName, fileName)
}
