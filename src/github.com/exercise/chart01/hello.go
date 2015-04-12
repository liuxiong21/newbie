package chart01

import (
	"fmt"
	"os"
	"strings"
)

func ShowHello() {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello,", who)
}
