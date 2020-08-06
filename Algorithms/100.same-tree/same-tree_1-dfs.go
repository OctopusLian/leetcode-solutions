/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//解法1，深度优先搜索
    if p == nil && q == nil {
        //如果两个二叉树都为空，则两个二叉树相同。
        return true
    }
    if p == nil || q == nil {
        //如果两个二叉树中有且只有一个为空，则两个二叉树一定不相同。
        return false
    }
    if p.Val != q.Val {
        //如果两个二叉树都不为空，那么首先判断它们的根节点的值是否相同，若不相同则两个二叉树一定不同
        return false
    }
    //若相同，再分别判断两个二叉树的左子树是否相同以及右子树是否相同。
    //这是一个递归的过程，因此可以使用深度优先搜索，递归地判断两个二叉树是否相同。
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}