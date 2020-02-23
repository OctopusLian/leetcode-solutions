/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    var pre *ListNode
    cur := head  //声明一个*ListNode类型变量并将head值赋给cur
    for cur != nil{
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}
