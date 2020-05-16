/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    tail, i := head, 0
    for ;tail != nil && i < k; i++ {
        tail = tail.Next
    }
    if i == k {
        tail = reverseKGroup(tail, k)
        for j := 0; j < k; j++ {
            tail, head.Next, head = head, tail, head.Next
        }
        head = tail
    }
    return head
}