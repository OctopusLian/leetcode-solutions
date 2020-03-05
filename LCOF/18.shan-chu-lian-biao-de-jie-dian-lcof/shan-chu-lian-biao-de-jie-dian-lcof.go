/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Val: 0, Next: head} // 假头
    first := dummy // 双指针1
    second := dummy.Next // 双指针2
    for {
        if second.Val == val {
            first.Next = second.Next  //删除指定节点
            break
        }
        first = first.Next
        second = second.Next
    }

    return dummy.Next
}
