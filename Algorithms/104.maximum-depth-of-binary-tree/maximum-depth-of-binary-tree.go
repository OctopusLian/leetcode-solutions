/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } else {
        left := maxDepth(root.Left) // 左子树的最大深度
        right := maxDepth(root.Right) // 右子树的最大深度
        return int(math.Max(float64(left), float64(right)) + 1) // 深度加上根节点
    }
}
