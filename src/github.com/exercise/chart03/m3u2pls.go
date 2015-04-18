package main

import (
	"os"
	"strings"
	"log"
	"io"
	"fmt"
	"path/filepath"
	"bufio"
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

func M3u2pls(source,target string){
	if strings.TrimSpace(source)==""{
		log.Fatal("Source file path is invalid")
	}
	
	sourceFile,err := os.Open(source)
	defer sourceFile.Close()
	if err!=nil{
		panic(err)
	}
	err = os.MkdirAll(filepath.Dir(target),0777)
	if err!=nil{
		panic(err)
	}
	targetFile,err := os.Create(target)
	if err!=nil{
		panic(err)
	}
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(targetFile)
	defer writer.Flush()
	for {
		line,err := reader.ReadBytes('\n')
		if err!=nil{
			if err== io.EOF{
				break
			}
			panic(err)
		}
		lineStr := string(line[:len(line)-1])
		if lineStr=="#EXTM3U"{
			writer.WriteString("[playlist]\n")
		}
		fmt.Println(lineStr)
	}
}



