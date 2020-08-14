/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func listOfDepth(tree *TreeNode) []*ListNode {
    if tree == nil {
		return nil
	}
	m := []*TreeNode{tree}
	var list []*ListNode

	for len(m) > 0 {

		var temp []*TreeNode
		node:= new(ListNode)
        tempnode :=node
		for i := 0; i < len(m); i++ {
			tempnode.Next = &ListNode{
				Val: m[i].Val,
			}
 	    	tempnode = tempnode.Next           
			if m[i].Left != nil {
				temp = append(temp, m[i].Left)
			}
			if m[i].Right != nil {
				temp = append(temp, m[i].Right)
			}
		}
		m = temp
		list = append(list, node.Next)
	}
	return list
}