package solution

import "fmt"

func Solution() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, &TreeNode{5, nil, nil}},
		Right: &TreeNode{3, nil, &TreeNode{6, nil, nil}}}
	ret := rightSideView(root)
	fmt.Println(ret)
}
