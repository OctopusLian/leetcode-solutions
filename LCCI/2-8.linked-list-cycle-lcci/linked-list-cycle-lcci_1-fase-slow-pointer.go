/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    //解法1，快慢指针法，然后相遇后再从链表头开始
    if head == nil {
        return nil
    }
    fast, slow := head, head
    for {
        if slow.Next == nil {
            return nil
        }
 
        slow = slow.Next
        if fast.Next == nil || fast.Next.Next == nil {
            return nil
        }
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }

    third := head
    for slow != third && slow != nil && third != nil {
        slow = slow.Next
        third = third.Next
    }
    return slow
}