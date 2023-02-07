package main

import (
	"fmt"
	"sort"
	"strings"
)

func gridChallenge(grid []string) string {
	// Write your code here
	for _, g := range grid {
		chars := strings.Split(g, "")
		sort.Strings(chars)
		fmt.Printf("%v - %v\n", g, chars)
	}

	return ""
}

func main() {
	fmt.Println(gridChallenge([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"}))
}
