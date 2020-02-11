/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    arr := []int{}  //建立一个空数组
    cur := head
    for cur != nil {  //循环将链表里的值写入数组中
        arr = append(arr,cur.Val)
        cur = cur.Next
    }

    left := 0
    right := len(arr)-1

    for left <= right {  //循环判断这个数组是不是回文数组
        if arr[left] != arr[right] {
            return false
        }
        left++
        right--
    }
    return true
}
