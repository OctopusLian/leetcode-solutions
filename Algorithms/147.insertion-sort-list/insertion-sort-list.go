/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{Next: head}
    lastSorted, curr := head, head.Next
    for curr != nil {
        if lastSorted.Val <= curr.Val {
            lastSorted = lastSorted.Next
        } else {
            prev := dummyHead
            for prev.Next.Val <= curr.Val {
                prev = prev.Next
            }
            lastSorted.Next = curr.Next
            curr.Next = prev.Next
            prev.Next = curr
        }
        curr = lastSorted.Next
    }
    return dummyHead.Next
}
