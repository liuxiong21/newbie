package main

import (
	"fmt"
	"unicode"
)

type LowerCaser interface {
	LowerCase()
}

type UpperCaser interface {
	UpperCase()
}

type LowerUpperCaser interface {
	LowerCaser
	UpperCaser
}

type FixCaser interface {
	FixCase()
}

type ChangeCaser interface {
	LowerUpperCaser
	FixCaser
}

func changePartCase(changeCase ChangeCaser) {
	fmt.Println(changeCase)
	changeCase.LowerCase()
	fmt.Println(changeCase)
	changeCase.UpperCase()
	fmt.Println(changeCase)
	changeCase.FixCase()
	fmt.Println(changeCase)

}

func (part *Part) FixCase() {
	var chars []rune
	upper := true
	for _, c := range part.Name {
		if upper {
			chars = append(chars, unicode.ToUpper(c))
		} else {
			chars = append(chars, unicode.ToLower(c))
		}
		if unicode.IsSpace(c) || unicode.Is(unicode.Hyphen, c) {
			upper = true
		} else {
			upper = false
		}
	}
	part.Name = string(chars)
}
