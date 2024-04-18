package main

import (
	"fmt"
	"testing"
)

//输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	ShowPreOrder(root)

}

// 前序遍历
func ShowPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	ShowPreOrder(root.Left)
	ShowPreOrder(root.Right)
}

// 中序遍历
func ShowInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	ShowPreOrder(root.Left)
	fmt.Println(root.Val)
	ShowPreOrder(root.Right)
}

// 中序遍历
func ShowAfterOrder(root *TreeNode) {
	if root == nil {
		return
	}
	ShowPreOrder(root.Left)
	ShowPreOrder(root.Right)
	fmt.Println(root.Val)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
