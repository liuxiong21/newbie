package chart01 

import (
	"fmt"
	"log"
	"os"
)

func BigDigital(numStr string) {
	if len(numStr) == 0 {
		log.Fatal("Call argments too less");
		os.Exit(-1);
	}
	fmt.Println("Start...")
	var bigDigits = [][]string {
		{" 000 ","0   0", "0   0", "0   0", "0   0", "0   0", " 000 "},
		{" 1 ", "1 1", " 1 ", " 1 ", " 1 ", " 1 ", "111"}, 
		{" 222 ", "2   2", "   2  ", "  2 ", " 2 ", "2 ","22222"},
		{" 333", "3   3", "    3", "   33", "    3", "3   3","33333"},
		{" 444 ", "4 4", " 4 ", " 4 ", " 4 ", " 4 ","44444"},
		{" 555 ", "5 5", " 5 ", " 5 ", " 5 ", " 5 ","55555"},
		{" 666 ", "6 6", " 6 ", " 6 ", " 6 ", " 6 ","66666"},
		{" 777 ", "7 7", " 7 ", " 7 ", " 7 ", " 7 ","77777"},
		{" 888 ", "8 8", " 8 ", " 8 ", " 8 ", " 8 ","88888"},
	    {" 9999", "9      9", "9 9", " 9999", " 99", " 9"," 9999"},
	}
	for i:=0;i < len(bigDigits[0]);i++{
		for column := range numStr{
			var num = numStr[column]-'0'
			if num >=0 && num <=9{
				fmt.Print(bigDigits[num][i])
			}else{
				log.Fatal("Invalid digit input");
			}
		}
		fmt.Println("")
	}
	fmt.Println("End....")
}

