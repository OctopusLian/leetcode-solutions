/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    // 快慢指针 - 第一个分号是变量定义，第二个分号内容是循环条件，第三个分号是每一次遍历要做的事情
    for front := head; front != nil && front.Next != nil; front, head = front.Next.Next, head.Next{}
    return head
}
