package main

import (
	"fmt"
	"strings"
)

func typeAssertion() {
	var i interface{} = 100
	var s interface{} = [2]string{"China", "America"}
	if result, ok := i.(int); ok {
		fmt.Printf("%T,%v,%T\n", result, result, i)
	}
	fmt.Printf("%T\n", s)
	ss := s.([2]string)
	fmt.Printf("%T=%v\n", ss, ss)
}

func shadow() (err error) {
	x, err := check1()
	if err != nil {
		return err
	}
	if y, err := check2(); err != nil {
		fmt.Println(y)
		return err
	}
	fmt.Println(x)
	return err
}

func Suffix(filename string)(result string){
	sp := "."
	filename = strings.ToLower(filename)
	index := strings.LastIndex(filename, sp)
	var suffix string
	defer func(){
		if strings.Index(result,".")==0{
			result = result[1:]
		}
	}()
	if index > 0 {
		suffix = filename[index+1:]
		return Suffix(filename[:index])+sp+suffix
	}
	return ""
}

func check1() (count int, err error) {
	return
}

func check2() (count int, err error) {
	return 1, fmt.Errorf("%s", "found error")
}
