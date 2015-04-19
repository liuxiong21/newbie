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
}

func swapAndProduct(a,b,result *int){
	tmp := *a
	*a = *b
	*b = tmp
	*result = (*a) * (*b)
}

