/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	//解法1，层次遍历
	if root == nil {
        return nil
    }
    q := []*Node{root}
    for len(q) > 0 {
        tmp := q
        q = nil
        for i, node := range tmp {
            if i+1 < len(tmp) {
                node.Next = tmp[i+1]
            }
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
    return root
}