func findKthLargest(nums []int, k int) int {
    //k个pointer组成link
    head := createLink(k)
    for i := 0; i < len(nums); i++ {
        insert(head, nums[i])
    }
    return lastElem(head)
}

type Elem struct {
    n int
    next *Elem
}

func createLink(k int) *Elem {
    head := &Elem{n: math.MinInt32}
    dummyhead := head
    for i:=1; i < k; i++ {
        tmp := &Elem{n: math.MinInt32}
        head.next = tmp
        head = tmp
    }
    return dummyhead
}

func insert(head *Elem, n int) {
    for ;head != nil; {
        if n >= head.n {
            break
        }
        head = head.next
    }
    flush := n
    for ;head != nil; {
        head.n, flush = flush, head.n
        if flush == math.MinInt32 {
            break
        }
        head = head.next
    }
}

func lastElem(head *Elem) int {
    for ;head.next != nil; head = head.next{
    }
    return head.n
}
