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
    //1，将链表里的值顺序append进事先声明好的re数组；
    for ;head != nil; {
        re = append(re,head.Val)
        head = head.Next
    }
    //2，将re数组中的值倒序append进er数组后，返回。
    for i:=len(re)-1;i>=0;i-- {
        er = append(er,re[i])
    }
    return er
}
