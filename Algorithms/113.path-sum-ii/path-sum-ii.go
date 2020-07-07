/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) [][]int {
	 //解法1，递归
    ret := [][]int{}
    path(root, []int{}, sum, &ret)
    return ret
}
func path(root *TreeNode, sol []int, sum int, ret *[][]int){
    if root == nil {
        return 
    }
    if root.Left == nil && root.Right == nil && root.Val == sum {
        sol = append(sol, root.Val)
        *ret = append(*ret, append([]int(nil), sol...))
    }
    sol = append(sol, root.Val)
    path(root.Left, sol, sum-root.Val, ret) 
    path(root.Right, sol, sum-root.Val, ret)
}