package Week03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}

	parent := map[int]*TreeNode{}//存储父节点
	visited := map[int]bool{}//存储要

	dfs(root, parent)

	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}

	return nil
}

func dfs(r *TreeNode, parent map[int]*TreeNode) {
	if r == nil {
		return
	}
	if r.Left != nil {
		parent[r.Left.Val] = r
		dfs(r.Left, parent)
	}
	if r.Right != nil {
		parent[r.Right.Val] = r
		dfs(r.Right, parent)
	}
}
