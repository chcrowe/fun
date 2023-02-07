package main

import "fmt"

func getChildIndex(a []int, parentIndex int, boundary int) int {
	childIndex1 := 2*parentIndex + 1
	childIndex2 := 2*parentIndex + 2

	if childIndex1 >= boundary {
		return childIndex1
	} else if childIndex2 >= boundary {
		return childIndex1
	} else if a[childIndex1] > a[childIndex2] {
		return childIndex1
	} else {
		return childIndex2
	}
}

func swap(a []int, i1 int, i2 int) {
	temp := a[i1]
	a[i1] = a[i2]
	a[i2] = temp
}

func bubbleDown(a []int, parentIndex int, boundary int) {
	childIndex := getChildIndex(a, parentIndex, boundary)

	for childIndex < boundary && a[parentIndex] < a[childIndex] {
		swap(a, parentIndex, childIndex)
		parentIndex = childIndex
		childIndex = getChildIndex(a, parentIndex, boundary)
	}
}

func heapsort(a []int) []int {
	for i := len(a) - 1; -1 < i; i-- {
		bubbleDown(a, i, len(a))
		//fmt.Printf("\t%v\n", a)
	}

	for wall := len(a) - 1; -1 < wall; wall-- {
		swap(a, 0, wall)
		//fmt.Printf("**\t%v\n", a)
		bubbleDown(a, 0, wall)
		//fmt.Printf("\t%v\n", a)
	}

	return a
}

func main() {
	input := []int{4, 15, 16, 50, 8, 23, 42, 108}
	fmt.Println("before")
	fmt.Printf("%v\n", input)
	heapsort(input)
	fmt.Println()
	fmt.Println("after")
	fmt.Printf("%v\n", input)

	//4,8,15,16,23,42,50,108
}
