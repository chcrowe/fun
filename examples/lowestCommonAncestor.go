// Given the root node of a binary tree and two distinct values, find the lowest common ancestor.

/*
       (5) root
      /   \
     (2)  (7)
         /  \
      (4)   (8)
               \
               (9)

input = root, p=4, q=9
output = 7

create and store a external ancestor value

implement a recursive traversal routine
function prototype contain a treenode, p, q return true of false
  return true if either p or q is found

    if p and q are found (2 or greater) then the currentnode is the lowest common ancestor



*/

package main

import (
	"fmt"
)

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

type Solution struct {
	Ancestor *TreeNode
	root     *TreeNode
}

func (t *Solution) insert(data int) *Solution {
	if t.root == nil {
		t.root = &TreeNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (t *Solution) recurse(node *TreeNode, p *TreeNode, q *TreeNode) bool {

	if nil == node {
		return false
	}

	leftRes := 0
	leftOk := t.recurse(node.left, p, q)
	if true == leftOk {
		leftRes = 1
	}

	rightRes := 0
	rightOk := t.recurse(node.right, p, q)
	if true == rightOk {
		rightRes = 1
	}

	currentRes := 0
	if p.data == node.data || q.data == node.data {
		currentRes = 1
	}

	if currentRes+leftRes+rightRes >= 2 {
		t.Ancestor = node
	}

	return currentRes+leftRes+rightRes > 0
}

func (t *Solution) findLowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	t.recurse(root, p, q)
	return t.Ancestor
}

func (t *Solution) InsertAfter(node *TreeNode, index int, data int) bool {
	if node == nil {
		return false
	}

	if node.data == index {
		node.insert(data)
		return true
	}

	if true == t.InsertAfter(node.left, index, data) {
		return true
	}
	if true == t.InsertAfter(node.right, index, data) {
		return true
	}
	return false
}

func listPrint(node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("%c:%d\n", ch, node.data)
	listPrint(node.left, ns+2, 'L')
	listPrint(node.right, ns+2, 'R')

}

func main() {

	tree := &Solution{}
	tree.insert(5)
	tree.insert(2)
	tree.insert(7)

	//tree.insert(4)
	tree.InsertAfter(tree.root, 7, 4)

	tree.insert(8)
	tree.insert(9)

	p := &TreeNode{data: 4, left: nil, right: nil}
	q := &TreeNode{data: 9, left: nil, right: nil}
	tree.recurse(tree.root, p, q)

	fmt.Println()
	fmt.Println("tree contents...")
	listPrint(tree.root, 0, 'M')

	fmt.Println()
	fmt.Printf("Ancestor: %d\n", tree.Ancestor.data)
	fmt.Println()
}
