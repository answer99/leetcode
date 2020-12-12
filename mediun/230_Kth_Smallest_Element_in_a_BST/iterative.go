package main

import (
	"fmt"
)

func main() {
	r := []int{3, 1, 4, 0, 2}
	t := Tree{}
	// initialize tree
	for _, v := range r {
		t.insert(v)
	}

	// get the answer
	fmt.Println(kthSmallest(t.root, 1))
}

func kthSmallest(root *TreeNode, k int) int {
	var s []*TreeNode

	for {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		// pop
		result := s[len(s)-1]
		s = s[0 : len(s)-1]

		k = k - 1
		if k == 0 {
			return result.Val
		}
		root = result.Right
	}
}
