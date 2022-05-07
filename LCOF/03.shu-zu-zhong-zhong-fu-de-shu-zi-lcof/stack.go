/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    //构建一个栈，用来存储链表中每个节点的值
    stack := make([]int,0)
    //构建一个指针，指向链表的头结点位置，从它开始向后遍历
    curNode := head
    //不断的遍历原链表中的每个节点，直到为nil
    for curNode != nil {
        //把每个节点的值加入到栈中
        stack = append(stack,curNode.Val)
        //curNode向后移动
        curNode = curNode.Next
    }
    //获取栈的长度
    size := len(stack)
    //定义一个同样长度的数组 res
    res := make([]int,size)
    //遍历栈，从栈顶挨个弹出每个值，把这些值依次加入到数组res中
    for i := 0;i < size;i++ {
        //数组接收栈顶元素值
        res[i] = stack[len(stack)-i-1]
    }
    
    return res
}