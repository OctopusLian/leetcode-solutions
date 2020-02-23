/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder)==0{
        return nil
	}
	a:=new(TreeNode)
	for i,k:=range inorder{
		if k==preorder[0]{
			a.Val=k
			a.Left=buildTree(preorder[1:i+1],inorder[0:i])
			a.Right=buildTree(preorder[i+1:],inorder[i+1:])
			break
		}
	}
	return a
}
