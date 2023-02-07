package main

import (
	"fmt"
)

func lonelyinteger(a []int32) int32 {
	m := make(map[string]int32)

	for _, b := range a {
		m[fmt.Sprintf("%d", b)] += 1
	}

	var res int32 = -1

	for key, val := range m {
		//fmt.Printf("Key: %s, Value: %d\n", key, val)
		if 1 == val {
			fmt.Sscanf(key, "%d", &res)
			return res
		}
	}
	return -1
}

func main() {

	fmt.Printf("%d", lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1}))
}
