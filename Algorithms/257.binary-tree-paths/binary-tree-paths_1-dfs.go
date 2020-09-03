/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var paths []string

func binaryTreePaths(root *TreeNode) []string {
	//解法1，深度优先搜索
	paths = []string{}
	constructPaths(root, "")
	return paths
 }
 
func constructPaths(root *TreeNode, path string) {
	 if root != nil {
		 pathSB := path
		 pathSB += strconv.Itoa(root.Val)
		 if root.Left == nil && root.Right == nil {
			 paths = append(paths, pathSB)
		 } else {
			 pathSB += "->"
			 constructPaths(root.Left, pathSB)
			 constructPaths(root.Right, pathSB)
		 }
	 }
}