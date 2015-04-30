package main

import (
	"fmt"
	"strings"
	"unicode"
	//"bytes"
	"io"
	"log"
)

type RuneForRuneFunc func(rune) rune

var removePunctuation RuneForRuneFunc = func(c rune) rune {
	if unicode.Is(unicode.Terminal_Punctuation, c) {
		return -1
	}
	return c
}

func processing(phrases []string) {
	for index, pharse := range phrases {
		phrases[index] = strings.Map(removePunctuation, pharse)
	}
}

type Counter int

func (counter *Counter) Increment() {
	*counter++
}

func (counter *Counter) Decrement() {
	*counter--
}

func (counter Counter) IsZero() bool {
	return counter == 0
}

type Part struct {
	Id   int
	Name string
}

func (part *Part) ToLowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) ToUpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

func (part Part) String() string {
	return fmt.Sprintf("<<%d %q>>", part.Id, part.Name)
}

type Item struct {
	id       string
	price    float64
	quantity int
}

func (item *Item) Cost() float64 {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	CatelogId int
}

func (item *SpecialItem) Cost() float64 {
	return item.Item.Cost() * float64(item.CatelogId)
}

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first  string
	second string
}

func (pair *StringPair) Exchange() {
	log.Println("entry...")
	pair.first, pair.second = pair.second, pair.first
}

func (pair *StringPair) Read(data []byte) (n int, err error) {
	log.Println("entry...")
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	var count int
	if pair.first != ""{
		n := copy(data,pair.first)
		pair.first = pair.first[n:]
		count += n;
	}
	if pair.second !=""{
		n := copy(data[count:],pair.second)
		pair.second = pair.second[n:]
		count += n;
	}
	/**if pair.first != "" {
		for _, bt := range []byte(pair.first) {
			data = append(data, bt)
			fmt.Println("len=", len(data))
		}
	}
	pair.first = ""

	if pair.second != "" {
		for _, bt := range []byte(pair.second) {
			data = append(data, bt)
			fmt.Println("len=", len(data))
		}
	}
	pair.second = ""
	fmt.Printf("%q\n",data)**/
	return count, nil
}

func (pair StringPair) String() string {
	return fmt.Sprintf("StringPair[%q,%q]", pair.first, pair.second)
}

type Point [2]int

func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

func (point Point) String() string {
	return fmt.Sprintf("Point[%d,%d]", point[0], point[1])
}
