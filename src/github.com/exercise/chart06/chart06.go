package main

import (
	"fmt"
	"io"
)

func main() {
	var phrases []string = []string{"Day; dusk, and light.", "All day long"}
	processing(phrases)
	fmt.Printf("%v\n", phrases)
	var counter Counter
	counter.Increment()
	fmt.Println(counter)
	counter.Increment()
	fmt.Println(counter)
	fmt.Println(counter.IsZero())
	part := Part{11, "liuxiong"}
	fmt.Println(part)
	item := Item{"001", 10.9, 10}
	fmt.Println(item.Cost())
	specialItem := SpecialItem{Item{"001", 10.9, 10}, 5}
	fmt.Println(specialItem.Cost())
	costLiteral := (*SpecialItem).Cost
	fmt.Println(costLiteral(&specialItem))
	stringPair := StringPair{"This is First", "This is Second"}
	fmt.Println(stringPair)
	exchange(&stringPair)
	fmt.Println(stringPair)
	datas := make([]byte, 100)
	var reader io.Reader = &stringPair
	size, err := reader.Read(datas)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%q\n", datas[:size])
	}
	fmt.Println(stringPair)
	bytes := make([]byte, 0)
	bytes = append(bytes, 'd')
	fmt.Printf("%q\n", bytes)
	part2 := Part{12, "thE case-insensible"}
	changePartCase(&part2)
	printPerson()
	printAuthor1()
	printAuthor2()
	printOptions()
	printFuzzyBool()
}

func exchange(function Exchanger) {
	function.Exchange()
}


