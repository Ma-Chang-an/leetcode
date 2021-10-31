package solution

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	que := []*TreeNode{root}
	return getLevelOrderBottom(que)
}

func getLevelOrderBottom(que []*TreeNode) [][]int {
	if len(que) == 0 {
		return [][]int{}
	}
	var curRet []int
	var nextLevel []*TreeNode
	for _, node := range que {
		curRet = append(curRet, node.Val)
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}
	return append(getLevelOrderBottom(nextLevel), curRet)
}
