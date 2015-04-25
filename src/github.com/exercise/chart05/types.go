package main

import (
	"fmt"
)

func typeAssertion() {
	var i interface{} = 100
	if result, ok := i.(int); ok {
		fmt.Printf("%T,%v,%T", result, result, i)
	}
}

func shadow()(err error){
	x,err:= check1()
	if err!=nil{
		return err
	}
	if y,err := check2();err!=nil{
		fmt.Println(y)
		return err
	}
	fmt.Println(x)
	return err
}

func check1() (count int,err error){
	return;
}


func check2() (count int,err error){
	return 1,fmt.Errorf("%s","found error");
}