package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 两个节点之间的距离便是两个节点分别到达他们最近共同祖先的距离之和
func findDistance(root *TreeNode, p int, q int) int {
	if p == q {
		return 0
	}
	ancestor := lowestCommonAncestorIn(root, p, q)
	dis2p := distanceAncestor2children(ancestor, p)
	dis2q := distanceAncestor2children(ancestor, q)
	return dis2p + dis2q
}
func lowestCommonAncestorIn(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p || root.Val == q {
		return root
	}
	leftCommonAncestor := lowestCommonAncestorIn(root.Left, p, q)
	rightCommonAncestor := lowestCommonAncestorIn(root.Right, p, q)
	if leftCommonAncestor != nil && rightCommonAncestor != nil {
		return root
	}
	if leftCommonAncestor == nil {
		return rightCommonAncestor
	}
	return leftCommonAncestor
}
func distanceAncestor2children(ancestor *TreeNode, childrenVal int) int {
	if ancestor == nil {
		return -1
	}
	if ancestor.Val == childrenVal {
		return 0
	}
	disLeft := distanceAncestor2children(ancestor.Left, childrenVal)
	if disLeft != -1 {
		return disLeft + 1
	}
	disRight := distanceAncestor2children(ancestor.Right, childrenVal)
	if disRight != -1 {
		return disRight + 1
	}
	return -1
}
