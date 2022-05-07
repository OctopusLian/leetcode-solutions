/**
 * Definition for ListNode
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
    /*
     * @param head: the head of linked list.
     * @param val: An integer.
     * @return: a linked node or null.
     */
    public ListNode findNode(ListNode head, int val) {
        for (ListNode node = head;node != null;node = node.next) {
            if (node.val == val) {
                return node;
            }
        }
        return null;
    }
}