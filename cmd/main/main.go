package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alfiehiscox/jgc-vis/pkg/parser"
)

func main() {

	bs, err := os.ReadFile("logs/java8-test.log")
	if err != nil {
		panic(err)
	}

	fs := string(bs)
	ss := strings.Split(fs, "\n")

	for i, l := range ss {
		parser := parser.NewParser(l)
		log, err := parser.Parse()
		if err != nil {
			panic(err)
		}
		fmt.Println(i, log)
	}

}
