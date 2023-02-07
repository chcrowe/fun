package main

import "fmt"

func towerBreakers(n int32, m int32) int32 {
	if 1 == m || 0 == n%2 {
		return 2
	} else {
		return 1
	}
}

func main() {
	fmt.Println("a: ", towerBreakers(2, 6))
}
