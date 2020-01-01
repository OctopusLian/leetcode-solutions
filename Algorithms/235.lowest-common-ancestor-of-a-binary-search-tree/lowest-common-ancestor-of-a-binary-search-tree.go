/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     //如果左节点大于又节点，调换值
     if p.Val > q.Val {
        p,q = q,p
    }
    if root == nil || p.Val<=root.Val && root.Val<=q.Val {
        return root
    } else if q.Val<=root.Val {
        return lowestCommonAncestor(root.Left,p,q)
    } else {
        return lowestCommonAncestor(root.Right,p,q)
    }
}
