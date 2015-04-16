package main 

import (
	"fmt"
	"strings"
)

type BitFlag int

const(
	Active BitFlag = 1<<iota
	Send
	Receive
)

func (flag BitFlag) String() string{
	var flags []string
	if flag & Active == Active {
		flags = append(flags,"Active")
	}
	
	if flag & Send == Send {
		flags = append(flags,"Send")
	}
	
	if flag & Receive == Receive {
		flags = append(flags,"Receive")
	}
	if len(flags)>0{
		return fmt.Sprintf("Int(%d)(%s)",int(flag),strings.Join(flags,"|"))
	}
	return "Int()()"
}

func main() {
	flag := Active|Send
	fmt.Println(flag,BitFlag(0))
	TypeNum()
}

