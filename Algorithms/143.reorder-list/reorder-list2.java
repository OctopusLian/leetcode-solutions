/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public void reorderList(ListNode head) {
        if (head == null || head.next == null) return;

        // find middle
        ListNode slow = head, fast = head.next;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        ListNode rHead = slow.next;
        slow.next = null;

        // reverse ListNode on the right side
        ListNode prev = null;
        while (rHead != null) {
            ListNode temp = rHead.next;
            rHead.next = prev;
            prev = rHead;
            rHead = temp;
        }

        // merge two list
        rHead = prev;
        ListNode lHead = head;
        while (lHead != null && rHead != null) {
            ListNode temp1 = lHead.next;
            lHead.next = rHead;
            rHead = rHead.next;
            lHead.next.next = temp1;
            lHead = temp1;
        }
    }
}
