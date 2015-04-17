package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "中国china"
	fmt.Println(str[0])
	fmt.Println(str[3])
	fmt.Println(str[6])
	æs := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æs += string(char)
	}
	text := "中 国人 I am chinese"
	firstCharEndIndex := strings.Index(text," ")
	firstChar := text[:firstCharEndIndex]
	fmt.Println()
	fmt.Println(firstChar)
}
