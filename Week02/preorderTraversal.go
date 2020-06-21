package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var r []int
	var s []*TreeNode
	c := root
	for c != nil || len(s) != 0 {
		if c != nil {
			s = append(s, c)
			r = append(r, c.Val)
			c = c.Left
		} else {
			p := s[len(s)-1]
			s = s[:len(s)-1]
			c = p.Right

		}

	}
	return r

}
