/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
    if t2 == nil {
        return true
    } else if t1 != nil && t2 != nil {
        return isSame(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2) 
    }
    return false
}

func isSame(p, q *TreeNode) bool {
    if q == nil {
        return true
    } else if p == nil {
        return false
    }
    return p.Val == q.Val && isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}