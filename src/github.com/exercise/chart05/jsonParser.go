package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
	Salary  float64
}

func (user User) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "User[name=%s,age=%d,address=%s,salary=%.2f]", user.Name, user.Age, user.Address, user.Salary)
	return buffer.String()
}

func JsonParser(jsonStr string) {
	var object interface{}
	if err := json.Unmarshal([]byte(jsonStr), &object); err != nil {
		fmt.Println(err)
	} else {
		propertyMap := object.(map[string]interface{})
		for key, value := range propertyMap {
			fmt.Printf("%s=%#v\n", key, value)
		}
	}
	var user User 
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}
