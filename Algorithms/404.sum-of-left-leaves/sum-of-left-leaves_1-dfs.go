/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isLeafNode(node *TreeNode) bool {
    return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
    if node.Left != nil {
        if isLeafNode(node.Left) {
            ans += node.Left.Val
        } else {
            ans += dfs(node.Left)
        }
    }
    if node.Right != nil && !isLeafNode(node.Right) {
        ans += dfs(node.Right)
    }
    return
}

func sumOfLeftLeaves(root *TreeNode) int {
	//解法1，深度优先搜索
    if root == nil {
        return 0
    }
    return dfs(root)
}