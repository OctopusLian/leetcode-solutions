/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if p.Val >= root.Val {
        return inorderSuccessor(root.Right, p)
    } else {
        left := inorderSuccessor(root.Left, p)
        if left == nil {
            return root
        }
        return left
    }
}