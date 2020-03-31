func sortArray(nums []int) []int {
    return heapSort(nums)
}

//堆排序
func heapSort( arr []int ) []int {
    n := len(arr)
	for i := (n-1)/2; i >= 0; i-- {
		shiftDown( arr, n, i )
	}

	for i := n-1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown( arr, i, 0 )
	}
    return arr
}

func shiftDown( arr []int, n, k int )  {
	e := arr[k]
	for 2*k+1 < n {
		j := 2*k+1
		if j + 1 < n && arr[j] < arr[j+1] {
			j = j + 1
		}
		if e >= arr[j] {
			break
		}
		arr[k] = arr[j]
		k = j
	}
	arr[k] = e
}
