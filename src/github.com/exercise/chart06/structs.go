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
