package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unicode"
	"io"
	"os"
	"strconv"
	"regexp"
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
	stringMap()
	readRune()
	stringConv()
	regexpMatch()
	M3u2pls("/opt/book/goeg/src/m3u2pls/David-Bowie-Singles.m3u","/opt/tmp/go/David-Bowie-Singles.pls")
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

func stringMap(){
	var mapFunc = func(c rune)rune{
		if c>127{
			return -1
		}
		return c
	}
	fmt.Println(strings.Map(mapFunc,"this is test 顶顶顶顶 fff"))
	
}

func readRune(){
	reader := strings.NewReader("中国ch")
	for{
		char,size,err := reader.ReadRune()
		if err!=nil{
			if err== io.EOF{
				break
			}
			panic(err)
		}
		fmt.Fprintf(os.Stdout,"%U '%c' %#T %d % X\n",char,char,char,size,[]byte(string(char)))
	}
}

func stringConv(){
	var buffer []byte
	bools := []string{"true","false","1","0","True"}
	for _,b := range bools{
		if result,err := strconv.ParseBool(b);err==nil{
			buffer = strconv.AppendBool(buffer,result)
			buffer = append(buffer,' ')
		}
	}
	fmt.Print(string(buffer))
	fmt.Println("\n",unicode.SimpleFold('中'))
}

func regexpMatch(){
	myreg :=  regexp.MustCompile(`\s*([[:alpha:]]\w*)\s*:\s*(\w+)`)
	lines := []string{" name : freeman","email:liuxiong21@gmail.com","sex:man age:19"}
	for _,line := range lines{
		if matches := myreg.FindAllStringSubmatch(line,-1);len(matches)>0{
			for _,match := range matches{
				fmt.Println(match[1],"-->",match[2])
			}
		}
	}
}
