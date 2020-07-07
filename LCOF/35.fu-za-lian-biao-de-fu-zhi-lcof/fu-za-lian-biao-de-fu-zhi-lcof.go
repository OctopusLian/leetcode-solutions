/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 func dfs(head *Node,used map[*Node]*Node)*Node{
    if head==nil{
        return nil
    }
    
    if c,ok:=used[head];ok{
        return c
    }

    res:=&Node{head.Val,nil,nil}
    used[head] = res
    res.Next = dfs(head.Next,used)
    res.Random = dfs(head.Random,used)
    return res
    
}

func copyRandomList(head *Node) *Node {
	//解法1，递归
    used:=make(map[*Node]*Node,0)
    return dfs(head,used)
}