package main

import (
	"fmt"
)


type Person struct{
	Title string
	Forenames []string
	Surname string
}

func (person Person) String() string{
	return fmt.Sprintf("Title=[%q],Forenames=%q,Surname=[%q]",person.Title,person.Forenames,person.Surname)
}

type Author1 struct{
	Names Person
	Titles []string
	YearBorn int
}

func (author Author1) String() string{
	return fmt.Sprintf("%v,Title=%q,Born=[%v]",author.Names,author.Titles,author.YearBorn)
}

type Author2 struct{
	Person
	Title string
	YearBorn int
}

func (author Author2) String() string{
	return fmt.Sprintf("Person.Title=[%q],Forenames=%q,Title=[%q],Born=[%v]",author.Person.Title,author.Forenames,author.Title,author.YearBorn)
}

func printPerson(){
	person1 := Person{"President",[]string{"zhuxi","zongtong"},"xidada"}
	fmt.Println(person1)
}

func printAuthor1(){
	person1 := Person{"President",[]string{"zhuxi","zongtong"},"xidada"}
	author1 := Author1{person1,[]string{"book01","book02"},1950} 
	fmt.Println(author1)	
}

func printAuthor2(){
	person1 := Person{"President",[]string{"zhuxi","zongtong"},"xidada"}
	author2 := Author2{person1,"this is title",1950} 
	fmt.Println(author2)	
}

type Optioner interface {
	Name() string
	IsValid() bool
}

type OptionCommon struct {
	ShortName string "this Short Name"
	LongName  string "this is Long Name"
}

type IntOption struct {
	OptionCommon
	Value, Max, Min int
}

func (option IntOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option IntOption) IsValid() bool {
	return option.Value >= option.Min && option.Value <= option.Max
}

type FloatOption struct {
	Optioner
	Value float64
}

type StringOption struct{
	Optioner
	Value string
}

func (option FloatOption) Name() string {
	return fmt.Sprintf("FloatOption[%v]", option.Value)
}

func (option FloatOption) IsValid() bool {
	return true
}

type GenericOption struct {
	OptionCommon
}

func (option GenericOption) Name() string {
	return name(option.ShortName, option.LongName)
}

func (option GenericOption) IsValid() bool {
	return true
}

func printOptions(){
	fileOption := StringOption{GenericOption{OptionCommon{"My name is ShortName","My name is LongName"}},
		"index.html"}
	floatOption := FloatOption{Optioner:GenericOption{OptionCommon{"s1","l1"}},Value:19.9}
	optioners := []Optioner{&fileOption,&floatOption}
	for _,op := range optioners{
		fmt.Println(op.Name(),"---",op.IsValid())
	}
}

func name(shortName, longName string) string {
	if shortName == "" {
		return shortName
	}
	return longName
}
