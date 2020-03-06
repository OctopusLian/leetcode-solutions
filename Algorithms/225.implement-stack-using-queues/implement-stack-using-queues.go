type MyStack struct {
    Queue1 []int
    Queue2 []int
    TopElement int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    var myStack MyStack
    return myStack
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Queue1 = append(this.Queue1,x)  //向队列Queue1里push一个x
    this.TopElement = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    length1 := len(this.Queue1)
    for i := 0; i < length1 - 1; i++ {
        // 取出每个元素
        item := this.Queue1[0]
        this.TopElement = item
        // 删除元素
        this.Queue1 = this.Queue1[1:]
        // 入队列 2
        this.Queue2 = append(this.Queue2, item)
    }
    target := this.Queue1[0]
    // 交换
    this.Queue1 = this.Queue2
    this.Queue2 = make([]int, 0)
    return target
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.TopElement
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.Queue1) == 0  //置空
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
