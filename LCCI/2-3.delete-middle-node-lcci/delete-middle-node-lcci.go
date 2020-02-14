/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    next := node.Next
    node.Val = next.Val
    node.Next = next.Next
}
