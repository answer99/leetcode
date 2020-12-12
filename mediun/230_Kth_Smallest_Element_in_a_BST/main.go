package main

import "fmt"

// Tree root
type Tree struct {
	root *TreeNode
}

// TreeNode a BST node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &TreeNode{Val: data}
	} else {
		t.root.insert(data)
	}
}

func (n *TreeNode) insert(data int) {
	if data == 0 {
		return
	}
	if data >= n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
		} else {
			n.Right.insert(data)
		}
	} else {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
		} else {
			n.Left.insert(data)
		}
	}
}

func traverseInOrder(n *TreeNode) []int {
	result := []int{}
	if n.Left != nil {
		result = append(result, traverseInOrder(n.Left)...)
	}
	result = append(result, n.Val)
	if n.Right != nil {
		result = append(result, traverseInOrder(n.Right)...)
	}
	return result
}

func traversePreOrder(n *TreeNode) []int {
	result := []int{}
	result = append(result, n.Val)
	if n.Left != nil {
		result = append(result, traverseInOrder(n.Left)...)
	}
	if n.Right != nil {
		result = append(result, traverseInOrder(n.Right)...)
	}
	return result
}

func kthSmallest(root *TreeNode, k int) int {
	inOrder := traverseInOrder(root)
	return inOrder[k-1]
}

func main() {
	r := []int{3, 1, 4, 0, 2}
	t := Tree{}
	for _, v := range r {
		t.insert(v)
	}
	fmt.Println(kthSmallest(t.root, 1))

}
