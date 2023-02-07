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
	matrix   [][]rune
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

func (t *Solution) InsertAfter(node *TreeNode, index int, data int) bool {
	if node == nil {
		return false
	}

	if node.data == index {
		node.insert(data)
		return true
	}

	if true == t.InsertAfter(node.left, index, data) || true == t.InsertAfter(node.right, index, data) {
		return true
	}

	return false
}

func (t *Solution) listPrint(root *TreeNode, node *TreeNode, mid int, depth int, ch rune) {
	if node == nil {
		return
	}

	/*
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ total width
	                      M                     center
	                    L    R                 center - 2, center + 2
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _009_ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _ _ 009 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ 009 _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _009_ _099_ _ _ _ _ _ _ _ _ _ _ _ _009_ _099_ _ _ _ _ _
	 _ _ _ _ _ _ 009 009 099 099 _ _ _ _ _ _ _ _ _ 009 _ _ _ _ 099 _ _ _ _
	 _ _ _ _ _009_ _ _ _ _ _ _ _099_ _ _ _ _ _ _099_ _ _ _ _ _ _ _099_ _ _
	 _ _ _ 099 _ _ _ _ _ _ _ _ _ _ 099 _ _ _ 099 _ _ _ _ _ _ _ _ _ _ 099 _
	 _ _099_ _ _ _ _ _ _ _ _ _ _ _ _ _099_099_ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
	 _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _

	*/

	for i := 0; i < mid+(depth*4); i++ {
		fmt.Printf(" ")
	}

	fmt.Printf("(%c%d)\n", ch, node.data)
	depth++

	t.listPrint(root, node.left, mid, depth, 'L')
	t.listPrint(root, node.right, mid, depth, 'R')

}

func printMatrix(matrix [][]rune) {

	for i := 0; i < len(matrix); i++ {
		fmt.Printf(string(matrix[i]))
		fmt.Println()
	}
}

func NewResultsMatrix(length, width int) [][]rune {
	result := make([][]rune, length)
	for i := 0; i < length; i++ {
		result[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			ch := '_'
			if 0 == j%2 {
				ch = ' '
			}
			result[i][j] = ch
		}
	}
	return result
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
	fmt.Println("5 2 7 4 8 9")
	fmt.Println("tree contents...")
	tree.listPrint(tree.root, tree.root, 16, 0, 'M')

	fmt.Println()
	fmt.Printf("Ancestor of 4 & 9: %d\n", tree.Ancestor.data)
	fmt.Println()

	printMatrix(tree.matrix)
	fmt.Println()
	fmt.Println()

}
