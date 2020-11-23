/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func merge(head1, head2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
    if head == nil {
        return head
    }

    if head.Next == tail {
        head.Next = nil
        return head
    }

    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
            fast = fast.Next
        }
    }

    mid := slow
    return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
    return sort(head, nil)
}