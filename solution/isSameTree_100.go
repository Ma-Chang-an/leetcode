package solution

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeBFS(p, q)
}

func isSameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	leftOk := isSameTree(p.Left, q.Left)
	if !leftOk {
		return false
	}
	return isSameTree(p.Right, q.Right)
}

func isSameTreeBFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	queP := []*TreeNode{p.Left, p.Right}
	queQ := []*TreeNode{q.Left, q.Right}
	for len(queP) > 0 && len(queQ) > 0 {
		curP := queP[0]
		curQ := queQ[0]
		queP = queP[1:]
		queQ = queQ[1:]
		if curP == nil && curQ == nil {
			continue
		}
		if curP == nil || curQ == nil || curP.Val != curQ.Val {
			return false
		}
		queP = append(queP, curP.Left, curP.Right)
		queQ = append(queQ, curQ.Left, curQ.Right)
	}
	return true
}
