package main

import (
	"fmt"
	"strings"
)

func findDuplicates(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		index := (nums[i] - 1) % lenNums

		nums[index] = lenNums + nums[index]
	}

	k := 0

	for i := 0; i < lenNums; i++ {
		if nums[i] > 2*lenNums {
			nums[k] = i + 1
			k++
		}
	}

	return nums[0:k]

}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func on_the_hunt(s string, arr []string, csv string) string {
	// Write your code here
	sChars := []rune(s)
	csvChars := strings.Split(csv, ",")

	m := make(map[string]int)

	for _, ch := range sChars {
		m[string(ch)] += 1
	}

	for _, ch := range csvChars {
		m[string(ch)] += 1
	}

	for _, ch := range arr {
		m[string(ch)] += 1
	}

	var results []string

	for key, val := range m {
		//fmt.Printf("Key: %s, Value: %d\n", key, val)
		if 1 < val {
			results = append(results, key)
		}
	}

	return strings.Join(results, "")
}

func main() {

	// b := byte(13)
	// bits := fmt.Sprintf("%08b", b) // Output: Go
	// fmt.Println(bits)
	// fmt.Println(reverse(bits))

	// output := findDuplicates([]int{1, 2, 3, 2, 4, 3})
	// fmt.Println(output)

	s := on_the_hunt("abcde", []string{"d", "a", "f", "g", "e"}, "d,a,g,e,b")
	fmt.Println(s)
}
