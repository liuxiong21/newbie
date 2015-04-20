package main

import (
)

func swapAndProduct(a,b,result *int){
	tmp := *a
	*a = *b
	*b = tmp
	*result = (*a) * (*b)
}

