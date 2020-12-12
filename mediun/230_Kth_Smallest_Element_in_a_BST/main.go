package main

import "fmt"

func main() {
	r := []int{3, 1, 4, 0, 2}
	t := Tree{}
	for _, v := range r {
		t.insert(v)
	}
	fmt.Println(kthSmallest(t.root, 1))
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
