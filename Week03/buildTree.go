package Week03


func buildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil || len(preorder) != len(inorder) || len(preorder) == 0{
		return nil
	}
	root := &TreeNode{preorder[0],nil,nil}
	i := 0
	for  ;i < len(inorder);i++{
		if preorder[i] == inorder[i]{//找到对应的根节点
			break
		}
	}

	root.Left = buildTree(preorder[1:len(preorder[:i])+1],inorder[:i])
	root.Right = buildTree(preorder[len(preorder[:i])+1:],inorder[i:])

	return  root


}
