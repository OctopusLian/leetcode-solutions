/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//解法一，深度优先搜索
type Codec struct {
    l []string
}

func Constructor() Codec {
    return Codec{}    
}

func rserialize(root *TreeNode, str string) string {
    if root == nil {
        str += "null,"
    } else {
        str += strconv.Itoa(root.Val) + ","
        str = rserialize(root.Left, str)
        str = rserialize(root.Right, str)
    }
    return str
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    return rserialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    l := strings.Split(data, ",")
    for i := 0; i < len(l); i++ {
        if l[i] != "" {
            this.l = append(this.l, l[i])
        }
    }
    return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
    if this.l[0] == "null" {
        this.l = this.l[1:]
        return nil
    }

    val, _ := strconv.Atoi(this.l[0])
    root := &TreeNode{Val: val}
    this.l = this.l[1:]
    root.Left = this.rdeserialize()
    root.Right = this.rdeserialize()
    return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */