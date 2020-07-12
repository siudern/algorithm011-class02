package Week_02

//给定一个二叉树，返回它的 前序 遍历。
//
// 示例:
//
//输入: [1,null,2,3]
//   1
//     \
//      2
//     /
//   3
//
//输出: [1,2,3]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 方法一，利用递归调用
var res []int

func preHelper(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val)
		preHelper(root.Left)
		preHelper(root.Right)
	}
}


func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	preHelper(root)
	return res
}


//方法二，利用栈
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		lastIndex := len(stack) - 1
		root = stack[lastIndex]
		stack = stack[:lastIndex]
	}
	return res
}
