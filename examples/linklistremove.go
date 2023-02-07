package main

import (
	"fmt"
)

type ListNode struct {
	next *ListNode
	data int64
}

type LinkedList struct {
	head     *ListNode
	tail     *ListNode
	prevTail *ListNode
}

func (t *LinkedList) update() {
	if nil == t.head {
		return
	}
	var temp *ListNode = t.head
	var prev *ListNode = nil
	var lastprev *ListNode = nil

	for nil != temp {
		lastprev = prev
		prev = temp
		temp = temp.next
	}

	t.prevTail = lastprev
	t.tail = prev
}

func (t *LinkedList) Add(data int64) *LinkedList {
	if t.head == nil {
		t.head = &ListNode{data: data, next: nil}
	} else {
		t.head.Add(data)
	}
	t.update()
	return t
}

func (t *LinkedList) Remove(i int) *ListNode {
	if nil == t.head {
		return nil
	}

	var temp *ListNode = t.head
	var prev *ListNode = nil
	delNode := &ListNode{data: 0, next: nil}

	for n := 0; nil != temp && n <= i; n++ {
		if n == i {
			delNode.data = temp.data
			prev.next = temp.next
			temp = nil
			break
		}
		prev = temp
		temp = temp.next
	}

	t.update()
	return delNode

}

func (t *LinkedList) RemoveTail() *ListNode {
	if nil == t.head {
		return nil
	}

	tailNode := &ListNode{data: t.tail.data, next: nil}

	t.tail = nil
	t.prevTail.next = nil
	t.prevTail = nil

	t.update()
	return tailNode

}

func (n *ListNode) Add(data int64) {
	if n == nil {
		return
	} else if n.next == nil {
		n.next = &ListNode{data: data, next: nil}
	} else {
		n.next.Add(data)
	}
}

func listPrint(node *ListNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.data)
	listPrint(node.next)
}

func main() {

	ll := &LinkedList{}
	ll.Add(9)
	ll.Add(9)
	ll.Add(12)
	ll.Add(1)
	ll.Add(3)
	ll.Add(5)
	ll.Add(7)
	ll.Add(5)
	ll.Add(3)
	ll.Add(1)
	ll.Add(6)

	fmt.Println()
	fmt.Println("initial list contents...")
	listPrint(ll.head)
	fmt.Printf(" (tail: %d)", ll.tail.data)
	fmt.Println()
	fmt.Println()

	fmt.Printf("remove index 2... (%d)", ll.Remove(2).data)
	fmt.Println()
	listPrint(ll.head)
	fmt.Printf(" (tail: %d)", ll.tail.data)
	fmt.Println()
	fmt.Println()

	fmt.Printf("remove last 2 items... ")
	for n := 0; n < 2; n++ {
		fmt.Printf("%d ", ll.RemoveTail().data)
	}
	fmt.Println()

	fmt.Println("udpated list contents...")
	listPrint(ll.head)
	fmt.Printf(" (tail: %d)", ll.tail.data)

	fmt.Println()
	fmt.Println()

	fmt.Println("add 3 nodes (17, 23, 28)")
	ll.Add(17)
	ll.Add(23)
	ll.Add(28)

	fmt.Println("final list contents...")
	listPrint(ll.head)
	fmt.Printf(" (tail: %d)", ll.tail.data)
	fmt.Println()
}
