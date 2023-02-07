package main

import (
	"fmt"
	"strconv"
	"strings"
)

var result []string

func permute(source string, substr string, depth int) {
	if depth == len(source) {
		result = append(result, substr)
		return
	}

	_, err := strconv.Atoi(string(source[depth]))
	if err == nil {
		permute(source, substr+string(source[depth]), depth+1)
	} else if string(source[depth]) == strings.ToUpper(string(source[depth])) {
		permute(source, substr+strings.ToLower(string(source[depth])), depth+1)
		permute(source, substr+string(source[depth]), depth+1)
	} else {
		permute(source, substr+strings.ToUpper(string(source[depth])), depth+1)
		permute(source, substr+string(source[depth]), depth+1)
	}
}

func main() {

	permute("A1bC2", "", int(0))

	fmt.Println(result)
}
