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
        //暴力法
        if(headA==null||headB==null){
            //若有一个链表为空，则不可能有相交
            return null;
        }
        ListNode curA=headA;
        ListNode curB=headB;
        for(;curA!=null;curA=curA.next){
            curB=headB;
            for(;curB!=null;curB=curB.next){
                if(curA==curB){
                    return curA;
                }
            }
        }
        return null;
    }
}
