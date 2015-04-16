package main

import (
	"fmt"
	"github.com/exercise/chart01"
)

func main() {
	fmt.Println("Begin")
	//callShowHello()
	callBigDigital()
	//callMyStack()
	//callAmericanise()
	//callPolar2Cartesian()
	fmt.Println("End")
}

func callShowHello(){
	chart01.ShowHello()
}

func callBigDigital(){
	chart01.BigDigital("123456")
}

func callMyStack(){
	var mystack chart01.Stacker
	mystack.Push("this is test")
	fmt.Println(mystack.Top())
	mystack.Push(0)
	fmt.Println(mystack.Top())
}

func callAmericanise(){
	chart01.Americanise()
}

func callPolar2Cartesian(){
	chart01.Polar2cartesian()
}