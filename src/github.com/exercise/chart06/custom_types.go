package main

import (
	//"fmt"
	"strings"
	"unicode"
)

type RuneForRuneFunc func(rune)rune

var removePunctuation RuneForRuneFunc = func(c rune) rune{
	if unicode.Is(unicode.Terminal_Punctuation,c){
		return -1
	}
	return c
}

func processing(phrases []string){
	for index,pharse := range phrases{
		phrases[index] = strings.Map(removePunctuation,pharse)
	}
}

type Counter int

func (counter *Counter) Increment(){
	*counter++;
}

func (counter *Counter) Decrement(){
	*counter--;
}

func (counter Counter) IsZero()bool{
	return counter==0
}

