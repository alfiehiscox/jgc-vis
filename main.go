package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alfiehiscox/jgc-vis/parser"
)

func main() {

	bs, err := os.ReadFile("logs/java8-test.log")
	if err != nil {
		panic(err)
	}

	fs := string(bs)
	ss := strings.Split(fs, "\n")

	s := ss[3]
	parser := parser.NewParser(s)
	log, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(log)
}
