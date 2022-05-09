/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	//特殊情况处理，删除的节点是头结点时
	if head.Val == val {
		return head.Next
	}
	//设置两个指针，一个指针指向当前的节点
	pre := head
	//一个指针指向当前节点的下一个节点
	cur := head.Next
	//当cur为空或cur节点值等于val时跳出循环
	for cur != nil && cur.Val != val {
		//两个指针不断的向前移动
		//pre来到cur的位置
		pre = cur
		//cur来到下一个节点位置
		cur = cur.Next
	}
	//删除指定节点
	pre.Next = cur.Next
	//最后返回链表头结点即可
	return head
}