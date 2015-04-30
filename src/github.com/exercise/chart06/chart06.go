package main 

import (
	"fmt"
)

func main() {
	var phrases []string = []string{"Day; dusk, and light.","All day long"}
	processing(phrases)
	fmt.Printf("%v\n",phrases)
	var counter Counter
	counter.Increment()
	fmt.Println(counter)
	counter.Increment()
	fmt.Println(counter)
	fmt.Println(counter.IsZero())
}


