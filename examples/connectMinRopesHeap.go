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

func removePeak(arr []int, heapBoundary int) int {
	peak := arr[0]
	swap(arr, 0, heapBoundary)
	bubbleDown(arr, 0, heapBoundary)
	return peak
}

func connectNRopesMinIntermediate(a []int) int {

	// add all to queue/min heap
	minHeapQ := make([]int, len(a))
	for i, val := range a {
		minHeapQ[i] = val
	}

	// sum ropes until 1 rope remaining
	firstMin, secondMin := 0, 0
	currentCost, totalCost := 0, 0
	for 1 < len(minHeapQ) {
		firstMin = minHeapQ[len(minHeapQ)-1]
		minHeapQ = minHeapQ[:len(minHeapQ)-1]
		secondMin = minHeapQ[len(minHeapQ)-1]
		minHeapQ = minHeapQ[:len(minHeapQ)-1]
		//sum the mins for current cost and add that to running total
		currentCost = firstMin + secondMin
		totalCost += currentCost
		// add current cost to min heap
		minHeapQ = append(minHeapQ, currentCost)

	}

	return totalCost
}

func connectNRopesMinAdvanced(a []int) []int {
	for i := len(a) - 1; -1 < i; i-- {
		bubbleDown(a, i, len(a))
		//fmt.Printf("\t%v\n", a)
	}

	heapSize := len(a)
	oldPeak := 0
	totalCost, currentCost := 0, 0

	for heapSize > 1 {
		heapSize--
		// remove peak and decrease heapsize
		oldPeak = removePeak(a, heapSize)
		// add old peak to new heap peak to get current min cost and new rope
		currentCost = a[0] + oldPeak
		// add current cost to total cost
		totalCost += currentCost
		// replace heap peak with new rope cost and bubble down
		a[0] = currentCost
		bubbleDown(a, 0, heapSize)
		//fmt.Printf("\t%v\n", a)
	}

	return a
}

func main() {
	input := []int{4, 3, 2, 6}
	fmt.Println("before")
	fmt.Printf("%v\n", input)
	total := connectNRopesMinIntermediate(input)
	fmt.Println(total)
	// fmt.Println("after")
	// fmt.Printf("%v\n", input)

	//4,8,15,16,23,42,50,108
}
