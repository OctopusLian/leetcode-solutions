/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	//解法1，模拟
    if root == nil {
        return &TreeNode{Val: val}
    }
    p := root
    for p != nil {
        if val < p.Val {
            if p.Left == nil {
                p.Left = &TreeNode{Val: val}
                break
            }
            p = p.Left
        } else {
            if p.Right == nil {
                p.Right = &TreeNode{Val: val}
                break
            }
            p = p.Right
        }
    }
    return root
}