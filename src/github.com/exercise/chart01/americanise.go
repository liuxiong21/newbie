package chart01

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"io"
	"path/filepath"
	"regexp"
	"strings"
)

var mapTableFile = "/opt/book/goeg/src/americanise/british-american.txt"

//frist arg:/opt/book/goeg/src/americanise/input.txt

func Americanise() error {
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	replaceFunc, err := makeReplaceFunction(mapTableFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var wordRx = regexp.MustCompile("[A-Za-z]+")
	var eof = false
	bReader := bufio.NewReader(inFile)
	bWriter := bufio.NewWriter(outFile)
	defer bWriter.Flush()
	for !eof {
		line,err := bReader.ReadString('\n')
		if err!=nil{
			if err==io.EOF{
				eof = true;
				err = nil
			}else{
				return err
			}	
		}
		line = wordRx.ReplaceAllStringFunc(line,replaceFunc)
		bWriter.WriteString(line)
	}
	return nil
}

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
	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, err
}
