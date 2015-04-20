package main

import (
	"fmt"
)

func main(){
	a := 100
	b := 1000
	product := 0 
	swapAndProduct(&a,&b,&product)
	fmt.Println(a,b,product)
	createArray()
	createSlice()
}

