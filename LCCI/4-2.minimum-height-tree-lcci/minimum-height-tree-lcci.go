/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    //递归
    if len(nums) == 0 {
        //考虑特殊情况
        return nil
    }
    if len(nums) == 1{
        //长度为1则返回数组第一个元素作为节点
        return &TreeNode{Val:nums[0]}
    }
    middle := len(nums) / 2
    head := &TreeNode{Val:nums[middle]} //将中间的元素取出作为头节点
    head.Left = sortedArrayToBST(nums[:middle])  //头结点左边的就是数组中间元素和左边的元素
    head.Right = sortedArrayToBST(nums[middle+1:])  //右边的就是右边的元素
    return head
}
