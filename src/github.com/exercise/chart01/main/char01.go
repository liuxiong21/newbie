package main

import (
	"fmt"
	"github.com/exercise/chart01"
)

func main() {
	fmt.Println("Begin")
	chart01.ShowHello()
	chart01.BigDigital("123456")
	fmt.Println("End")
	var mystack chart01.Stacker
	mystack.Push("this is test")
	fmt.Println(mystack.Top())
	mystack.Push(0)
	fmt.Println(mystack.Top())
}
