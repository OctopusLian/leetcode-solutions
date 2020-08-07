/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    var a,b = new(ListNode),new(ListNode)
    var p1,p2 = a,b
    for head != nil{
        if head.Val < x {
            p1.Next = &ListNode{head.Val, nil}
            p1 = p1.Next
        }else{
            p2.Next = &ListNode{head.Val, nil}
            p2 = p2.Next
        }

        head = head.Next
    }

    p1.Next = b.Next
    return a.Next
}