/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    temp := root.Left
    root.Left = mirrorTree(root.Right)  //左右交换
    root.Right = mirrorTree(temp)
    return root
}
