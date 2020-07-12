package Week_02


//给定一个二叉树，返回它的 后序 遍历。
//
//示例:
//
//输入: [1,null,2,3]
//   1
//     \
//      2
//     /
//   3
//
//输出: [3,2,1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//思路一：
//迭代

var postRes []int

func postHelper(root *TreeNode) {
	if root != nil {
		postHelper(root.Left)
		postHelper(root.Right)
		postRes = append(postRes, root.Val)
	}
}


func postorderTraversal(root *TreeNode) []int {
	postRes = make([]int, 0)
	postHelper(root)
	return postRes
}


//思路二：
//迭代
//后序遍历可以看做是前序遍历的变形版，先遍历右子树之后，再遍历左子树，最后reverse即可
//即  根-左-右 转变成 根-右-左，最后把左和根替换位置
func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Left)
			root = root.Right
		}
		lastIndex := len(stack) - 1
		root = stack[lastIndex]
		stack = stack[:lastIndex]
	}
	return reverse(res)
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}