/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	//解法1，队列
    if root==nil{
        return nil
    }
    var ret [][]int
    queue := []*TreeNode{root}
    //树顶视为第0层,从0开始
    level:=0
    for len(queue)!=0{
        temp := []*TreeNode{}
        ret = append(ret,make([]int,0))
        for _,v := range queue{
            if level & 1 ==0{
                //偶数层从左向右打印，数值追加至切片尾部
                ret[level]=append(ret[level],v.Val)
            }else{
                //奇数层从右向左打印，数值追加至切片头部
                s := []int{v.Val}
                ret[level]=append(s,ret[level]...)
            }
            //子树从左到右入队
            if v.Left!=nil{
                temp = append(temp,v.Left)
            }
            if v.Right!=nil{
                temp = append(temp,v.Right)
            }
        }
        //下一层
        queue = temp
        level++
    }
    return ret
}