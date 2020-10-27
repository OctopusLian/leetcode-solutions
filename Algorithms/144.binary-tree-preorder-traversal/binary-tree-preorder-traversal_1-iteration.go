/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode)(vals []int) {
	//解法1，迭代
    stack := []*TreeNode{}
    node := root
    for node != nil || len(stack) > 0 {
        for node != nil {
            vals = append(vals, node.Val)
            stack = append(stack, node)
            node = node.Left
        }
        node = stack[len(stack)-1].Right
        stack = stack[:len(stack)-1]
    }
    return
}