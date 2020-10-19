/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	//解法1，线性表
    if head == nil {
        return
    }
    nodes := []*ListNode{}
    for node := head; node != nil; node = node.Next {
        nodes = append(nodes, node)
    }
    i, j := 0, len(nodes)-1
    for i < j {
        nodes[i].Next = nodes[j]
        i++
        if i == j {
            break
        }
        nodes[j].Next = nodes[i]
        j--
    }
    nodes[i].Next = nil
}