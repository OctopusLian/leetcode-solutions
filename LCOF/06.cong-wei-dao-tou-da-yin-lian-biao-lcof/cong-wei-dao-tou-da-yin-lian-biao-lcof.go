/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    var re []int
    var er []int
    for ;head != nil; {
        re = append(re,head.Val)
        head = head.Next
    }
    for i:=len(re)-1;i>=0;i-- {
        er = append(er,re[i])
    }
    return er
}
