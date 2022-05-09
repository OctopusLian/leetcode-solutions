/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode deleteNode(ListNode head, int val) {
        //特殊情况处理，删除的节点是头结点时
        if (head.val == val) return head.next;
        //设置两个指针，一个指针指向当前的节点
        ListNode pre = head;
        //一个指针指向当前节点的下一个节点
        ListNode cur = head.next;
        //当cur为空或cur节点值等于val时跳出循环
        while (cur != null && cur.val != val) {
            //两个指针不断的向前移动
            //pre来到cur的位置
            pre = cur;
            //cur来到下一个节点位置
            cur = cur.next;
        }
        //相当于覆盖掉了cur的节点值
        pre.next = cur.next;
        //最后返回链表头结点即可
        return head;
    }
}