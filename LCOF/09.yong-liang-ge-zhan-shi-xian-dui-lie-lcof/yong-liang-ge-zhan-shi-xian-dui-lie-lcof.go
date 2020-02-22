type CQueue struct {
    //栈in负责输入，栈out负责从头输出
    in []int
    out []int
}


func Constructor() CQueue {
    return CQueue{
        //为两个栈预申请空间，减少内存消耗
        make([]int,0,5),
        make([]int,0,5),
    }
}


func (this *CQueue) AppendTail(value int)  {
    if value<1 && value>10000{
        return 
    }else{
        this.in=append(this.in[:],value)
    }
}


func (this *CQueue) DeleteHead() int {
    //情况1：无人进栈
    if len(this.in)==0{
        return -1
    }
    //情况2：出入相等
    if len(this.in)==len(this.out){
        return -1
    }
    //“out栈的长度”刚好等于“要输出的数在in栈中的位置”
    out:=this.in[len(this.out)]
    this.out=append(this.out[:],this.in[len(this.out)])
    return out
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
