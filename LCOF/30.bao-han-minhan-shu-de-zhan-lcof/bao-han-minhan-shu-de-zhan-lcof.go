type MinStack struct {
    elems []int //存放原始数据
    mins []int //存放最小值
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        elems: make([]int, 0),
        mins: make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    //第一次append，存放原始数据
    this.elems = append(this.elems, x)
    if len(this.mins) == 0 {
        this.mins = append(this.mins, x)
    } else {
        if this.mins[len(this.mins) - 1] >= x {
            this.mins = append(this.mins, x)
        }
    }
}


func (this *MinStack) Pop()  {
    elem := this.elems[len(this.elems) - 1]
    this.elems = this.elems[:len(this.elems)-1]
    if elem == this.mins[len(this.mins)-1] {
        this.mins = this.mins[:len(this.mins)-1]
    }
}


func (this *MinStack) Top() int {
    return this.elems[len(this.elems)-1]
}


func (this *MinStack) Min() int {
    return this.mins[len(this.mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
