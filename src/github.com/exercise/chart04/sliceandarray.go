package main

import (
	"fmt"
)

func createArray(){
	array1 := [10]int{}
	//array2 := [10]int{100,1000}
	array := [...]int{10,100,1000,1000}
	fmt.Println(len(array),len(array1))
	//array is not slice
	//array = append(array,100000)
}

func createSlice(){
	slice := []int{10}
	fmt.Println(len(slice),cap(slice))
	slice2 := make([]int,10,100)
	slice = append(slice,100)
	fmt.Println(len(slice),cap(slice))
	fmt.Println(len(slice2),cap(slice2))
}

