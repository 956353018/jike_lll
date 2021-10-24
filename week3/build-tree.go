package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (node *TreeNode) PreOrder(n *TreeNode) {
	if n != nil {
		fmt.Printf("%d ", n.Val)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var build func(l1 int, r1 int, l2 int, r2 int) *TreeNode

	//[l1,r1]是后序（左，右，根），[l2,r2]是中序（左、根、右）
	build = func(l1 int, r1 int, l2 int, r2 int) *TreeNode{
		//4.边界条件
		if l1 > r1 {
			return nil
		}
		//1.建立根
		root := &TreeNode{Val: postorder[r1]}

		//3.根据中序找根的位置mid
		mid := l2
		for {
			if inorder[mid] == root.Val {
				break
			}
			mid++
		}

		//2.构建根的左右子树
		//l2~mid-1左子树范围
		//mid+1,r2右子树范围
		leftTreeNum := (mid - 1) - l2 + 1
		//rightTreeNum := r2 - (mid + 1) + 1
		root.Left = build(l1, l1 + leftTreeNum - 1, l2, mid - 1)
		root.Right = build(l1 + leftTreeNum, r1 - 1, mid + 1, r2)

		return root
	}
	return build(0, len(postorder) - 1, 0, len(inorder) - 1)
}

func main() {
	inorder := []int{9,3,15,20,7}
	postorder := []int{9,15,7,20,3}
	root := buildTree(inorder, postorder)

	root.PreOrder(root)
	fmt.Printf("\n")
}
