package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateNode(val int) * TreeNode {
	return &TreeNode{val, nil,nil}
}

func (node *TreeNode) PreOrder(n *TreeNode) {
	if n != nil {
		fmt.Printf("%d ", n.Val)
		node.PreOrder(n.Left)
		node.PreOrder(n.Right)
	}
}

func dfsBST(cur *TreeNode, sum *int) {
	if cur != nil {
		//fmt.Printf("%d\n", cur.Val)
		//遍历右子树,一直到右子树底开始,里面的节点值都比当前的节点值大
		dfsBST(cur.Right, sum)

		//处理当前节点，sum加上当前节点，sum始终保存比当前节点大的值
		*sum += cur.Val
		cur.Val = *sum
		//fmt.Printf("%d\n", *sum)

		//遍历左子树
		dfsBST(cur.Left, sum)
	}
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	sum := 0
	dfsBST(root, &sum)
	return root

}

func main() {
	root := CreateNode(4)
	root.Left = CreateNode(1)
	root.Right = CreateNode(6)

	root.Left.Left = CreateNode(0)
	root.Left.Right = CreateNode(2)
	root.Left.Right.Right = CreateNode(3)

	root.Right.Left = CreateNode(5)
	root.Right.Right = CreateNode(7)
	root.Right.Right.Right = CreateNode(8)

	root.PreOrder(root)
	fmt.Printf("\n")
	newRoot := convertBST(root)
	newRoot.PreOrder(newRoot)
}