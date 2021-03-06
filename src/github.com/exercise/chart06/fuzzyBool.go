package main

import (
	"fmt"
)

func printFuzzyBool() {
	a,err := NewFuzzyBool("ddd")
	if err!=nil{
		fmt.Println(err)
	}
	b, _ := NewFuzzyBool(.25)
	c, _ := NewFuzzyBool(.75)
	d := c.Copy()
	if err := d.Set(1); err != nil {
		fmt.Println(err)
	}
	process(a, b, c, d)
	s := []*FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *FuzzyBool) {
	fmt.Println("Original:", a, b, c, d)
	fmt.Println("Not: ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(),
		d.Not().Not())
	fmt.Print("0.And(.25)→", a.And(b), "• .25.And(.75)→", b.And(c),
		"• .75.And(1)→", c.And(d), " • .25.And(.75,1)→", b.And(c, d), "\n")
	fmt.Print("0.Or(.25)→", a.Or(b), "• .25.Or(.75)→", b.Or(c),
		"• .75.Or(1)→", c.Or(d), " • .25.Or(.75,1)→", b.Or(c, d), "\n")
	fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), c.Less(a))
	fmt.Println("Bool: ", a.Bool(), b.Bool(), c.Bool(), d.Bool())
	fmt.Println("Float: ", a.Float(), b.Float(), c.Float(), d.Float())
}

type FuzzyBool struct {
	value float32
}

func (fuzzy *FuzzyBool) Set(val interface{}) error{
	fval, err := parseFloatForValue(val)
	if err!=nil{
		return err
	}
	fuzzy.value = fval
	return nil
}

func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

func NewFuzzyBool(val interface{}) (*FuzzyBool, error) {
	fval, err := parseFloatForValue(val)
	return &FuzzyBool{value: fval}, err
}

func parseFloatForValue(value interface{}) (val float32, err error) {
	switch value := value.(type) {
	case float32:
		val = float32(value)
	case float64:
		val = float32(value)
	case int:
		val = float32(value)
	case bool:
		val = 0
		if value {
			val = 1
		}
	default:
		return 0, fmt.Errorf("Found type[%T] nocompatibly,expect Boolean or Number", value)
	}
	return val, nil
}

func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*fuzzy.value)
}

func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

func (fuzzy *FuzzyBool) And(first *FuzzyBool,
	rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if minimum > other.value {
			minimum = other.value
		}
	}
	return &FuzzyBool{minimum}
}

func (fuzzy *FuzzyBool) Or(first *FuzzyBool,
	rest ...*FuzzyBool) *FuzzyBool {
	maximum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if maximum < other.value {
			maximum = other.value
		}
	}
	return &FuzzyBool{maximum}
}

func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value >= .5
}

func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}
