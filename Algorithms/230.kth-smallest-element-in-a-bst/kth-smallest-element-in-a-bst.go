/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    // 求左子树节点数
    leftCount := nodeCount(root.Left)
    // 根就是第k个
    if leftCount == k - 1 {
        return root.Val
    } else if leftCount >= k {
        return kthSmallest(root.Left, k)
    } else {
        return kthSmallest(root.Right, k - leftCount - 1)
    }
    return 0
}

func nodeCount(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + nodeCount(root.Left) + nodeCount(root.Right)
}
