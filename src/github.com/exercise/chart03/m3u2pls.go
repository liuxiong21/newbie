package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
	"strconv"
)

/**
pls format:
[playlist]
File1=Music/David Bowie/Singles 1/01-Space Oddity.ogg
Title1=David Bowie - Space Oddity
Length1=315
...
NumberOfEntries=33
Version=2

m3u format:
#EXTM3U
#EXTINF:315,David Bowie - Space Oddity
Music/David Bowie/Singles 1/01-Space Oddity.ogg

*/

type Song struct{
	title string
	filepath string
	length int
}

func M3u2pls(source, target string) {
	if strings.TrimSpace(source) == "" {
		log.Fatal("Source file path is invalid")
	}

	sourceFile, err := os.Open(source)
	defer sourceFile.Close()
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(filepath.Dir(target), 0777)
	if err != nil {
		panic(err)
	}
	targetFile, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(targetFile)
	defer writer.Flush()
	titleRex := regexp.MustCompile(`^#EXTINF\s*:\s*(?P<length>\d+),(?P<title>.+)`)
	var counter int
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lineStr := string(line[:len(line)-1])
		if lineStr == "#EXTM3U" {
			writer.WriteString("[playlist]\n")
		} else if titleRex.MatchString(lineStr) {
			song := Song{}
			groups := titleRex.FindAllStringSubmatch(lineStr, -1)
			for i := 0; i < len(groups); i++ {
				song.length,err = strconv.Atoi(groups[i][1])
				song.title = groups[i][2]
				line,err = reader.ReadBytes('\n')
				song.filepath = string(line[:len(line)-1])
				counter++
				writer.WriteString(fmt.Sprintf("File%d=%s\n",counter,song.filepath))
				writer.WriteString(fmt.Sprintf("Title%d=%s\n",counter,song.title))
				writer.WriteString(fmt.Sprintf("Length%d=%d\n",counter,song.length))
			}
		}
	}
	writer.WriteString(fmt.Sprintf("NumberOfEntries=%d\n",counter))
	writer.WriteString("Version=2")
}
