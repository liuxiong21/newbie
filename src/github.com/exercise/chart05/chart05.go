package main

import (
	"fmt"
)

func init(){
	fmt.Println("Chart05.go file init function called")
}

func main() {
	typeAssertion()
	fmt.Println(shadow())
	filepath := "/opt/tools/go.tar.gz"
	fmt.Printf("Suffix(%s)=%s\n", filepath, Suffix(filepath))
	fmt.Println("BoundedValue:",BoundedValue(10,100,50))
	TypeSwitch("this is string",30,int64(100),float64(1000.00),true)
	JsonParser(`{"name":"liuxiong21","age":18,"address":"guangzhou","salary":100.9}`)
	DoubleCounter()
	//SelectChannel()
	ErrorRecover()
	addZip := MakeAndAddSuffix(".zip")
	fmt.Println(addZip("maven"))
}
