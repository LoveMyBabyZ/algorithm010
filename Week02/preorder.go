package main

type Node struct {
	Val      int
	Children []*Node
}

//手动维护一个栈进行递归调用
type Stack struct {
	Val *Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	s := []*Node{root}
	for len(s) > 0 {
		p := s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, p.Val)
		for i := len(p.Children) - 1; i >= 0; i-- {
			if child := p.Children[i]; child != nil {
				s = append(s, child)
			}

		}
	}
	return res
}

//回调函数
func preorder2(root *Node) []int {
	res := []int{}
	return  helper(root,res)

}

func helper(root *Node ,res []int)  []int{
	if root == nil {
		return res
	}
	res = append(res,root.Val)
	for _,child := range root.Children{
		res = helper(child,res)
	}
	return res
}
