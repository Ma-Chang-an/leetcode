package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 寻找节点的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftCommonAncestor := lowestCommonAncestor(root.Left, p, q)
	if leftCommonAncestor == nil {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if leftCommonAncestor != p && leftCommonAncestor != q {
		return leftCommonAncestor
	}
	rightCommonAncestor := lowestCommonAncestor(root.Right, p, q)
	if rightCommonAncestor != nil {
		return root
	}
	return leftCommonAncestor
}
