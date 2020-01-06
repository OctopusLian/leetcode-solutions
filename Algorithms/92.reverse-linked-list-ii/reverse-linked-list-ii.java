/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode reverseBetween(ListNode head, int m, int n) {
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        
        // find the mth node
        ListNode premNode = dummy;
        for (int i = 1; i < m; i++) {
            premNode = premNode.next;
        }
        
        // reverse node between m and n
        ListNode prev = null, curr = premNode.next;
        while (curr != null && (m <= n)) {
            ListNode nextNode = curr.next;
            curr.next = prev;
            prev = curr;
            curr = nextNode;
            m++;
        }
        
        // join head and tail before m and after n
        premNode.next.next = curr;
        premNode.next = prev;
        
        return dummy.next;
    }
}
