package main

import (
	"strings"

	"github.com/golang/tour/wc"
)

func WordCount(s string) map[string]int {
	result := map[string]int{}
	f := strings.Fields(s)
	for _, v := range f {
		result[v] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
