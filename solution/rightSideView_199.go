package solution

//199. 二叉树的右视图
//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func rightSideView(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	que := []*TreeNode{root}
	for len(que) > 0 {
		ret = append(ret, que[len(que)-1].Val)
		var nextRow []*TreeNode
		for _, node := range que {
			if node.Left != nil {
				nextRow = append(nextRow, node.Left)
			}
			if node.Right != nil {
				nextRow = append(nextRow, node.Right)
			}
		}
		que = nextRow
	}
	return ret
}
