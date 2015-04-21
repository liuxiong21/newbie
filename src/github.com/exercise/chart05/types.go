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
