package Week_03

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
//注意:
//你可以假设树中没有重复的元素。
//
//例如，给出
//
//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//返回如下的二叉树：
//
//       3
//      / \
//     9  20
//       /  \
//      15   7
//
////来源：力扣（LeetCode）
////链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
////著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。



//思路：
// 任意一棵树：
//前序遍历为 [ 根节点，[左子树前序遍历结果], [右子树遍历结果]
//中序遍历为 [[左子树中序遍历结果], 根节点, [右子树中序遍历结果]]
//在中序遍历的切片中定位到根节点，就可以知道左右子树的节点数，通过在中序遍历的节点数定位前序遍历的左右子树长度，进行定位
//在中序遍历中定位根节点，根据题意可知元素非重复，则可利用hashmap的方式把每个树节点的位置记录下来，优化时间复杂度到O(1)

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
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
