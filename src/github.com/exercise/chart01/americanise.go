package chart01

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func filenamesFromCommandLine() (inFilename, outFilename string, err error) {
	fmt.Printf("Args len %d\n", len(os.Args))
	if len(os.Args) == 1 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
		err = fmt.Errorf("Usage:%s [<]inFile.txt [>]outFile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}

	if inFilename != "" && inFilename == outFilename {
		log.Fatal("Forbid override the inFile")
	}
	return inFilename, outFilename, err
}

func makeReplaceFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	content := string(rawBytes)
	usForBritish := make(map[string]string)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}
	fmt.Println(content)
	return func(word string) string{
		if usWord, found := usForBritish[word];found{
			return usWord
		}
		return word
	}, err
}
