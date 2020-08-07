/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, sum int) int {
    var m = 0
	slove(root,sum,&m)
	return m
}
func slove(root *TreeNode,sum int,m *int){
	if root==nil{
		return
	}
	helpDFSPathSum(root,sum,m)
	slove(root.Left,sum,m)
	slove(root.Right,sum,m)
}

func helpDFSPathSum( root *TreeNode,target int,m *int){
	if root==nil{
		return
	}
	target-=root.Val
	if target==0{
		*m++
	}
	helpDFSPathSum(root.Left,target,m)
	helpDFSPathSum(root.Right,target,m)
}