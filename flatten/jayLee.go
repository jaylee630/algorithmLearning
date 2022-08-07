package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 给你二叉树的根结点 root ，请你将它展开为一个单链表：
 * 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null。
 * 展开后的单链表应该与二叉树 先序遍历 顺序相同。
 */
func flatten(root *TreeNode) {
	// root = [1,2,5,3,4,null,6] -> [1,null,2,null,3,null,4,null,5,null,6]
	var (
		list     []*TreeNode
		preorder func(*TreeNode)
	)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		// 需要将每个节点按照前序顺序进行拼接，将左节点置空
		prev.Left, prev.Right = nil, curr
	}
	return
}
