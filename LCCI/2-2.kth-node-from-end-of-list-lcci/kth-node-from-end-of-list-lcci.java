/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public int kthToLast(ListNode head, int k) {
        ListNode p1 = head;
        ListNode p2 = head;
        for (int i = 0;i < k;i++) {
            if (p1 == null) return -1;
            p1 = p1.next;
        }
        while(p1 != null) {
            p1 = p1.next;
            p2 = p2.next;
        }
        return p2.val;
    }
}
