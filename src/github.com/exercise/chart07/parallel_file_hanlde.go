package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var w sync.WaitGroup

func parallelFileHandle() {
	filenameChan := sources("/opt/book")
	willHandledFiles := filterSuffixes([]string{".gz", ".pdf"}, filenameChan)
	//w.Add(1)
	finished := handleFile(willHandledFiles)
	<- finished
	//w.Wait()
}

func sources(dir string) <-chan string {
	files := getAllFiles("/opt/book/")
	result := make(chan string, 1000)
	go func() {
		for index, filename := range files {	
			fmt.Println(filename,index)		
			result <- filename
			time.Sleep(time.Duration(rand.Intn(10000)))
		}
		close(result)
	}()
	return result
}

func getAllFiles(dir string) []string {
	result := make([]string, 0, 1000)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result
}

func filterSuffixes(suffixes []string, filesChan <-chan string) <-chan string {
	result := make(chan string, 10)
	go func() {
		for file := range filesChan {
			result <- file
		}
		close(result)
	}()
	return result
}

func handleFile(files <-chan string) <-chan bool{
	result := make(chan bool)
	go func() {
		for file := range files {
			fmt.Println(file)
			time.Sleep(time.Second * time.Duration(1))
		}
		//w.Done()
		result <- true
	}()
	return result
}
