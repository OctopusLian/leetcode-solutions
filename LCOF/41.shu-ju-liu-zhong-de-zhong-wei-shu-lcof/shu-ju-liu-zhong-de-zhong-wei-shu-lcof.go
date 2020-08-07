type Heap []int

func (m *Heap)Swap(i,j int){
    (*m)[i],(*m)[j]=(*m)[j],(*m)[i]
}

func (m *Heap)Len()int{
    return len(*m)
}

func (m *Heap)Pop()(v interface{}){
    *m,v = (*m)[:len(*m)-1],(*m)[len(*m)-1]
    return
}

func (m *Heap)Push(v interface{}){
    *m = append(*m,v.(int))
}

type minHeap struct {
    Heap
}

func (m *minHeap)Less(i,j int)bool{
    return m.Heap[i] < m.Heap[j]
}

type maxHeap struct {
    Heap
}

func (m *maxHeap)Less(i,j int)bool{
    return m.Heap[i] > m.Heap[j]
}


type MedianFinder struct {
    RightMin *minHeap
    LeftMax  *maxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
    m := new(MedianFinder)
    m.RightMin = new(minHeap)
    m.LeftMax = new(maxHeap)
    heap.Init(m.LeftMax)
    heap.Init(m.RightMin)
    return *m
}


func (this *MedianFinder) AddNum(num int)  {
    // 压入数据时有两种情况；
    // （1）原有数据为偶数个，则有 leftheap.len() == rightheap.len()；
    //     此时将数据放入leftheap， 高明之处：需先将数据放在rightheap，再从rightheap中pop出一个元素，将其放入leftheap；
    //     上面做法的高明之处 在于 省去了判断 num 与 rightheap最小值 及 leftheap最大值 的比较，通过比较判断num应该放在哪个heap，
    //     然后再根据左右heap的长度，对左右heap进行调整
    //  (2) 原有数据为奇数个，则有 leftheap.len() == rightheap.len() + 1  (通过step 1 来保证当为奇数时，left总比right 大 1，这样可以简化判断)
    //     类似step 1 需先将数据放在leftheap，再从leftheap中pop出一个元素，将其放入rightheap中
    if this.LeftMax.Len() == this.RightMin.Len(){
        heap.Push(this.RightMin,num)
        heap.Push(this.LeftMax,heap.Pop(this.RightMin))
    }else{
        heap.Push(this.LeftMax,num)
        heap.Push(this.RightMin,heap.Pop(this.LeftMax))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if (this.RightMin.Len()+this.LeftMax.Len())%2==0{
        //此处根据heap的实现原理，直接获取heap的第一个元素，即为大顶堆的最大值
        return float64(this.LeftMax.Heap[0]+this.RightMin.Heap[0])/2
    }else{
        return float64(this.LeftMax.Heap[0])
    }
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */