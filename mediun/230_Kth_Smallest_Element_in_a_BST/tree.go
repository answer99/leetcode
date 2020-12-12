package main

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
