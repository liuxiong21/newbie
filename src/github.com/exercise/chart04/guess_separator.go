package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var nameMap = map[string]string{"\t": "tab-separator", " ": "whitespace", "|": "vertical-line", "-": "blash"}

func GuessSeparator(filepath string) []string {
	lines := readNLineFromFile(filepath, 5)
	separators := []string{"\t", " ", "|", "-"}
	statisMap := statis(lines, separators)
	fmt.Printf("%v\n", statisMap)
	reversion := reversionMap(statisMap)
	fmt.Printf("%v\n", reversion)
	var keys []int
	for key := range reversion {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	maxKey := keys[len(keys)-1]
	return reversion[maxKey]
}

func reversionMap(m map[string]int) map[int][]string {
	result := make(map[int][]string)
	for key, val := range m {
		vals, found := result[val]
		if !found {
			fmt.Printf("Key not found,%#v\n", vals)
		}
		vals = append(vals, key)
		result[val] = vals
	}
	return result
}

func statis(lines []string, separators []string) map[string]int {
	resultMap := make(map[string]int)
	for _, line := range lines {
		for _, sp := range separators {
			count := strings.Count(line, sp)
			namedKey := nameMap[sp]
			count += resultMap[namedKey]
			resultMap[namedKey] = count
		}
	}
	return resultMap
}

func readNLineFromFile(filepath string, n int) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lines[i] = string(bytes)
	}
	return lines

}
