package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")
	//callShowHello()
	//callBigDigital()
	//callMyStack()
	//callAmericanise()
	callPolar2Cartesian()
	fmt.Println("End")
}

func callShowHello(){
	ShowHello()
}

func callBigDigital(){
	BigDigital("123456")
}

func callMyStack(){
	var mystack Stacker
	mystack.Push("this is test")
	fmt.Println(mystack.Top())
	mystack.Push(0)
	fmt.Println(mystack.Top())
}

func callAmericanise(){
	Americanise()
}

func callPolar2Cartesian(){
	Polar2cartesian()
}