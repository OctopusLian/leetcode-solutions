func search(arr []int, target int) int {
	//解法1，深度优先搜索
    // 取巧，只可能发生在最开始的数组中
    if arr[0] == target {
        return 0
    }
    return dfs(arr, target, 0, len(arr)-1)
}

func dfs(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) >> 1
	for ; arr[mid] == target; mid-- {
		if arr[mid-1] != target {
			return mid
		}
	}

	if left == right {
		return -1
	}

    // 取巧，只可能发生在最开始的数组中，两边减1，递归处理
    if arr[mid] == arr[left] && arr[mid] == arr[right] {
        return dfs(arr, target, left+1, right-1)
    }

	if arr[mid] > target && (arr[left] > arr[mid] || arr[left] <= target) || (arr[right] < target && arr[left] > arr[mid]) {
		return dfs(arr, target, left, mid)
	}

	return dfs(arr, target, mid+1, right)
}