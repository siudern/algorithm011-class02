package Week_02

//给定一个二叉树，返回它的中序遍历。
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
//输出: [1,3,2]
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//思路1：
//利用递归调用的思想，中序遍历的顺序为 左-根-右
//时间复杂度：O(n)，递归函数 T(n)=2⋅T(n/2)+1
//空间复杂度：最坏情况下需要空间O(n)，平均情况为O(log n)。

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var treeSlice []int

func helper(root *TreeNode) {
	if root != nil {
		helper(root.Left)
		treeSlice = append(treeSlice, root.Val)
		helper(root.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	treeSlice = make([]int, 0)
	helper(root)
	return treeSlice
}



// 思路2：
//利用栈的特性
//时间复杂度：O(n)
//空间复杂度：O(n)
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {  // 当栈不为空的时候 或者存在根节点的时候则遍历
		for root != nil { // 不为空意味着存在子节点
			stack = append(stack, root) // 把子节点压入栈
			root = root.Left // 指定下一次继续往左子树查找
		} // 退出循环条件：即root.Left没有子节点了
		lastIndex := len(stack) -1  //记录最后一个节点的索引。顶栈
		res = append(res, stack[lastIndex].Val)  //把最后一个stack元素取出放入到res切片中。中序输出
		root = stack[lastIndex].Right  //再遍历右子树
		stack =stack[:lastIndex] //已经用过的栈元素要删除掉。出栈
	}
	return res
}