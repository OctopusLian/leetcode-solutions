/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    return (maxDepth(root) != -1)
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    if leftDepth == - 1 || rightDepth == -1 || (abs(leftDepth - rightDepth) > 1){
        return -1
    }
    return max(leftDepth,rightDepth) + 1
}

func abs(a int)int {
    if a < 0 {
        return a * (-1)
    } else {
        return a
    }
}

func max(a,b int)int {
    if a > b {
        return a
    } else {
        return b
    }
}