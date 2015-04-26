package main

import (
	"fmt"
)

func BoundedValue(mininum, value, maxinum int) int {
	switch {
	case value < mininum:
		return mininum
	case value > maxinum:
		return maxinum
	default:
		return value
	}
}

func ParseArchiveFile(filepath string) {
	switch suffix := Suffix(filepath); suffix {
	case "gz":
		fmt.Println("Gz file parser is executing")
	case "tar.gz":
		fallthrough
	case "tar":
		fmt.Println("Tar file parser is executing")
	}

	switch Suffix(filepath) {
	case "gz":
		fmt.Println("Gz file parser is executing")
	case "tar.gz", "tar":
		fmt.Println("Tar file parser is executing")
	}
}

func TypeSwitch(generics ...interface{}) {
	for _,item := range generics {
		switch item.(type) {
		case bool:
			fmt.Printf("Bool=%v\n", item)
		case int:
			fmt.Printf("Int=%v\n", item)
		case []int:
			fmt.Printf("[]Int=%v\n", item)
		case string:
			fmt.Printf("String=%v\n", item)
		default:
			fmt.Printf("Unknown Type=%v\n", item)
		}
	}
}
