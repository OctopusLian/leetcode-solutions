/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */

//超出时间限制
class Solution {
    public ListNode removeDuplicateNodes(ListNode head) {
        HashSet<Integer> set = new HashSet<Integer>();
        ListNode previous = null;
        while (head != null) {
            if (set.contains(head.val)) {
                previous.next = head.next;
            } else {
                set.add(head.val);
                previous = head;
            }
        }
        head = head.next;
        return head;
    }
}
