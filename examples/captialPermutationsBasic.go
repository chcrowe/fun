package main

import (
	"fmt"
	"strings"
)

var result []string

func permute(source string, substr string, depth int) {
	if depth == len(source) {
		result = append(result, substr)
		return
	}

	permute(source, substr+string(source[depth]), depth+1)
	permute(source, substr+strings.ToUpper(string(source[depth])), depth+1)
}

func main() {

	permute("abcde", "", int(0))

	fmt.Println(result)
}
