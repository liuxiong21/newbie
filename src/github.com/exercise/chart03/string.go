package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
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
	firstCharEndIndex := strings.Index(text, " ")
	firstChar := text[:firstCharEndIndex]
	fmt.Println()
	fmt.Println(firstChar)
	Humanize(10000000.299, 30, 10, "_","0")
}

func Humanize(amount float64, width, decimals int, sp,pad string) {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:]
	fractions := ""
	if decimals > 0 {
		fractions = fmt.Sprintf("%+.*f", decimals, cents)[2:]
	}
	for i := len(whole) - 3; i > 0; i = i - 3 {
		whole = whole[:i] + sp + whole[i:]
	}
	for i := 4; i < len(fractions); i = i + 3 {
		fractions = fractions[:i] + sp + fractions[i:]
		i++
	}
	number := whole + fractions
	if amount<0.0{
		number = "-"+number
	}
	gap := width - utf8.RuneCountInString(number)
	if gap>0{
		number = strings.Repeat(pad,gap)+number
	}
	fmt.Println(number)
}
