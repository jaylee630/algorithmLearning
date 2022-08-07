package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 */
func preorderTraversal(root *TreeNode) (val []int) {
	// root = [1,null,2,3] ->[1,2,3]
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) { // 当前节点的答案
		if node == nil {
			return
		}
		val = append(val, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
