func findKthLargest(nums []int, k int) int {
    //维护一个size为k的小顶堆，最后返回堆顶即可
    if k > len(nums)  {
        return 0
    }
    h := &MinHeap{}
    heap.Init(h)
    for i := 0; i < len(nums); i++ {
        heap.Push(h, nums[i])
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    return heap.Pop(h).(int)
}

type MinHeap []int
func (h MinHeap) Len() int {
    return len(h)
}
func (h MinHeap) Less(i, j int) bool {
    return  h[i] <= h[j]
}
func (h MinHeap) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}){
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := old.Len()-1
    x := old[n]
    *h = old[0:n]
    return x
}

