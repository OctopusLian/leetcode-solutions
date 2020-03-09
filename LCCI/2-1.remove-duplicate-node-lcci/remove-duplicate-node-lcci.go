/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
    //考虑特殊情况
    if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]bool, 0)
	fake := new(ListNode)
	fake.Next = head
	cur := head
	pre := fake

	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			cur = cur.Next
		} else {
			m[cur.Val] = true
            //删除重复的链表节点
			pre.Next = cur
			pre = pre.Next
			cur = cur.Next
		}
	}
	pre.Next = nil
	return head
}
