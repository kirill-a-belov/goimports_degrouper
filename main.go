package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

var (
	importRegExp    = regexp.MustCompile(`import *\(([\s\S]+?)\)`)
	emptyLineRegExp = regexp.MustCompile(`(?m)^$\n`)
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		panic("empty cmd param")
	}

	fileName := argsWithoutProg[0]
	if fileName == "" {
		panic("empty input file name")
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	result := importRegExp.ReplaceAllStringFunc(string(contents), removeLineBreakers)

	err = ioutil.WriteFile(fileName, []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}

func removeLineBreakers(input string) string {
	return emptyLineRegExp.ReplaceAllLiteralString(input,"")
}
