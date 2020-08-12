/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 func cloneGraph(node *Node) *Node {
	 //解法1，深度优先搜索
    visited := map[*Node]*Node{}
    var cg func(node *Node) *Node
    cg = func(node *Node) *Node {
        if node == nil {
            return node
        }

        // 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
        if _, ok := visited[node]; ok {
            return visited[node]
        }

        // 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
        cloneNode := &Node{node.Val, []*Node{}}
        // 哈希表存储
        visited[node] = cloneNode

        // 遍历该节点的邻居并更新克隆节点的邻居列表
        for _, n := range node.Neighbors {
            cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
        }
        return cloneNode
    }
    return cg(node)
}