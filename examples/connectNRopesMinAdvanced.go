package main

import "fmt"

func getChildIndex(a []int, parentIndex int, boundary int) int {
	childIndex1 := 2*parentIndex + 1
	childIndex2 := 2*parentIndex + 2

	if childIndex1 >= boundary {
		return childIndex1
	} else if childIndex2 >= boundary {
		return childIndex1
	} else if a[childIndex1] < a[childIndex2] {
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

	for childIndex < boundary && a[parentIndex] > a[childIndex] {
		swap(a, parentIndex, childIndex)
		parentIndex = childIndex
		childIndex = getChildIndex(a, parentIndex, boundary)
	}
}

func connectNRopesMinAdvanced(a []int) (int, []int) {
	for i := len(a) - 1; -1 < i; i-- {
		bubbleDown(a, i, len(a))
		fmt.Printf("\t%v\n", a)
	}

	heapSize := len(a)
	//oldPeak := 0
	totalCost, currentCost := 0, 0

	for heapSize > 1 {
		heapSize--
		// add old peak to new heap peak to get current min cost and new rope
		currentCost = a[0] + a[1]
		// add current cost to total cost
		totalCost += currentCost
		// replace heap peak with new rope cost and bubble down
		// fmt.Printf("\t%v\n", a)
		a = a[1:]
		a[0] = currentCost
		fmt.Printf("\t%v\n", a)
		bubbleDown(a, 0, heapSize)
		fmt.Printf("\t\t%v\n", a)
	}

	return totalCost, a
}

func main() {
	input := []int{4, 3, 2, 6}
	fmt.Println("before")
	fmt.Printf("%v\n", input)
	total, input := connectNRopesMinAdvanced(input)
	fmt.Println(total)
	fmt.Println("after")
	fmt.Printf("%v\n", input)

	//4,8,15,16,23,42,50,108
}
