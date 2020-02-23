/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
    n:=1    //本题从1开始计数
    cur:=head
    for head.Next!=nil{
        head=head.Next
        n++  //得到链表的长度
    }
    for i:=0;i<n-k;i++{
        cur=cur.Next //利用链表长度与k的差值找到倒数的节点
    }
    return cur
}
