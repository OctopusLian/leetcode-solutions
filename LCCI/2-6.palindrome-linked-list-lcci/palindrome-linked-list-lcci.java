/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public boolean isPalindrome(ListNode head) {
        ListNode reversed = reverseAndClone(head);  //第一步，反转整个链表
        return isEqual(head,reversed); //第二步，比较反转链表和原始链表，若两者相同，则该链表为回文
    }

    public ListNode reverseAndClone(ListNode node) {
        ListNode head = null;
        while(node != null) {
            ListNode n = new ListNode(node.val); //复制
            n.next = head;
            head = n;
            node = node.next;
        }
        return head;
    }

    public boolean isEqual(ListNode one,ListNode two) {
        while(one != null && two != null) {
            if (one.val != two.val) {
                return false;
            }
            one = one.next;
            two = two.next;
        }
        return one == null && two == null;
    }
}
