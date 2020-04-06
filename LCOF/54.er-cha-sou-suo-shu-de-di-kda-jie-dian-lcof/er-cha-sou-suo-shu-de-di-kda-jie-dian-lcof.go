/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
    count := make([]int, 0)
	dfskthLargest(root, &count)
	return count[k-1]
}

func dfskthLargest(root *TreeNode, count *[]int) {
	if root.Right != nil {
		dfskthLargest(root.Right, count)
	}
	if root != nil {
		*count = append(*count, root.Val)
	}
	if root.Left != nil {
		dfskthLargest(root.Left, count)
	}
}

