package Week_04

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
// 
//
//示例：
//二叉树：[3,9,20,null,null,15,7],
//
//	     3
//	    / \
//	   9  20
//	     /  \
//	    15   7
//返回其层次遍历结果：
//
//    [
//     [3],
//     [9,20],
//     [15,7]
//    ]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//记树上所有节点的个数为 n
//时间复杂度：每个点进队出队各一次，故渐进时间复杂度为 O(n)
//空间复杂度：队列中元素的个数不超过 n 个，故渐进空间复杂度为 O(n)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {   // 控制层级
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {    //遍历每一层的所有节点
			node := q[j]  //遍历节点
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p  // 把下一层层赋值给外层进行下一层的遍历，实现ret[index]刚好为层数
	}
	return ret
}