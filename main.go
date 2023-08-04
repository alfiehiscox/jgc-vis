package main

import (
	"fmt"

	"github.com/alfiehiscox/jgc-vis/parser"
)

func main() {
	s := "[GC (Allocation Failure) [PSYoungGen: 8704K->992K(9728K)] 8704K->3776K(31744K), 0.0073015 secs]"
	// s1 := "[Full GC (Ergonomics) [PSYoungGen: 1024K->0K(30720K)] [ParOldGen: 30716K->30648K(66560K)] 31740K->30648K(97280K), [Metaspace: 2859K->2859K(1056768K)], 0.0971264 secs]"
	tps := parser.Tokenize(s)

	for _, tp := range tps {
		fmt.Print(tp)
	}
}
