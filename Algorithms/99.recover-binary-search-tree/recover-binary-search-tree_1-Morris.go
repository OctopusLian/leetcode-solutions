/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
	//解法1，Morris 中序遍历
    var x, y, pred, predecessor *TreeNode

    for root != nil {
        if root.Left != nil {
            // predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
            predecessor = root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                predecessor = predecessor.Right
            }

            // 让 predecessor 的右指针指向 root，继续遍历左子树
            if predecessor.Right == nil {
                predecessor.Right = root
                root = root.Left
            } else { // 说明左子树已经访问完了，我们需要断开链接
                if pred != nil && root.Val < pred.Val {
                    y = root
                    if x == nil {
                        x = pred
                    }
                }
                pred = root
                predecessor.Right = nil
                root = root.Right
            }
        } else { // 如果没有左孩子，则直接访问右孩子
            if pred != nil && root.Val < pred.Val {
                y = root
                if x == nil {
                    x = pred
                }
            }
            pred = root
            root = root.Right
        }
    }
    x.Val, y.Val = y.Val, x.Val
}