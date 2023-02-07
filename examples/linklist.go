package main2

import (
	"fmt"
)

type ListNode struct {
	next *ListNode
	data int64
}

type LinkedList struct {
	root *ListNode
}

func (t *LinkedList) Add(data int64) *LinkedList {
	if t.root == nil {
		t.root = &ListNode{data: data, next: nil}
	} else {
		t.root.Add(data)
	}
	return t
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

func print(node *ListNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d ", node.data)
	print(node.next)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func recurse(s string, l int, r int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}

	if s[l] == s[r] {
		return 2 + recurse(s, l+1, r-1)
	} else {
		lRes := recurse(s, l+1, r)
		rRes := recurse(s, l, r-1)
		return Max(lRes, rRes)
	}
}

func main() {


	arr := 
	list := &LinkedList{}
	list.Add(9)
	list.Add(9)
	list.Add(1)
	list.Add(3)
	list.Add(5)
	list.Add(7)
	list.Add(5)
	list.Add(3)
	list.Add(1)
	list.Add(6)
	print(list.root)
}
