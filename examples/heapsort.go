package main

func getChildIndex(a []int, parentIndex int, boundary int) int {
	childIndex1 := 2*parentIndex + 1
	childIndex2 := 2*parentIndex + 2

}

func bubbleDown(a []int, parentIndex int, boundary int) {
	childIndex1 := getChildIndex(a, parentIndex, boundary)

}

func heapsort(a []int) []int {
	for i := len(a) - 1; -1 < i; i-- {
		bubbleDown(a, i, len(a))
	}
}

func main() {

}
