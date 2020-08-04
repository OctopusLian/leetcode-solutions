/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	//解法1，dfs	
    val := dfs(root)
    return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
    if node == nil {
        return []int{0, 0}
    }
    l, r := dfs(node.Left), dfs(node.Right)
    selected := node.Val + l[1] + r[1]
    notSelected := max(l[0], l[1]) + max(r[0], r[1])
    return []int{selected, notSelected}
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}