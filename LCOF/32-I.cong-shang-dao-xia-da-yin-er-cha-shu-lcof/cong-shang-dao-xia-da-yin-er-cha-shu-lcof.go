/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	//解法1，队列
    if root==nil{
        return nil
    }
    queue:=make([]*TreeNode,0,500)
    queue=append(queue,root)
    res:=make([]int,0,500)

    for len(queue)>0{
        node:=queue[0]
        if node.Left!=nil{
            queue=append(queue,node.Left)
        }
        if node.Right!=nil{
            queue=append(queue,node.Right)
        }

        res=append(res,node.Val)
        queue=queue[1:]
    }

    return res
}