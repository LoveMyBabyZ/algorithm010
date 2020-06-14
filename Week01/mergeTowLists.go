package main

type ListNode struct {
	Val  int
	Next *ListNode
}
//添加一个哨兵，然后以此 进行比较，将对应的节点信息放到哨兵的Next节点
//如果比较到最后还有剩余，则直接将剩下的信息放到Next节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	returnList := result
	for l1 != nil && l2 != nil {
		//如果小于
		if l1.Val < l2.Val {
			result.Next = l1
			l1 = l1.Next
		} else {
			result.Next = l2
			l2 = l2.Next
		}
		result = result.Next

	}
	if l1 == nil {
		result.Next = l2
	}
	if l2 == nil {
		result.Next = l1
	}
	return returnList.Next
}
