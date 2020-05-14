/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func convertBiNode(root *TreeNode) *TreeNode {
    head:=&TreeNode{0,nil,nil}
	cur:=head
	stack:=make([]*TreeNode,0)
	var node *TreeNode=root
	for node!=nil|| len(stack)!=0{
		if node!=nil{
			stack= append(stack, node)
			//先遍历左子树
            node=node.Left
		}else{
			//取出最新的入栈的元素
			node=stack[len(stack)-1]
			//取完之后栈要更新
			stack=stack[:len(stack)-1]
			//将node的左子树置为空
			node.Left=nil
			//将当前节点拼接到 cur上
			cur.Right=node
			//移动游标
			cur=node
			//遍历右子树
			node=node.Right
			
		}
	}
	return head.Right
}