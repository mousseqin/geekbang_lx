package main

func main() {

}

type TreeNode struct{}

// preorderTraversal 二叉树的前序遍历 :给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//func preorderTraversal(root *TreeNode) (vals []int) {
//	stack := []*TreeNode{}
//	node := root
//	for node != nil || len(stack) > 0 {
//		for node != nil {
//			vals = append(vals, node.Val)
//			stack = append(stack, node)
//			node = node.Left
//		}
//		node = stack[len(stack)-1].Right
//		stack = stack[:len(stack)-1]
//	}
//	return
//}
