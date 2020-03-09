/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
	//双指针法，消除长度差，最终两个指针要么相遇，要么都是空
        ListNode pa = headA;
        ListNode pb = headB;
        while (pa != pb) {
            if (pa != null) {
                pa = pa.next;
            }else{
                pa = headB;
            }
            if (pb != null) {
                pb = pb.next;
            }else{
                pb = headA;
            }
        }
        if (pa ==null) {
            return null;
        }
        return pa;
    }
}
