/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    //解法2，递归
    //递归终止条件是当前为空，或者下一个节点为空
    if head == nil || head.Next == nil {
        return head
    }
    //这里的cur就是最后一个节点
    cur := reverseList(head.Next)
    //如果链表是 1->2->3->4->5，那么此时的cur就是5
	//而head是4，head的下一个是5，下下一个是空
	//所以head.next.next 就是5->4
    head.Next.Next = head
    //防止链表循环，需要将head.next设置为空
    head.Next = nil
    //每层递归函数都返回cur，也就是最后一个节点
    return cur
}
