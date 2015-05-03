package main

import (
	"fmt"
	"github.com/exercise/chart06/omap"
	"github.com/user/stringutil"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	sortMap := omap.NewOmap(func(first, second interface{})bool {
			switch first.(type){
				case int,int32,int64:
					if first.(int64) < second.(int64){
						return true
					}
			}
			
			return false
	})
	fmt.Println("\n",sortMap.Len())
}
