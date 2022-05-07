/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public int[] reversePrint(ListNode head) {
        //构建一个栈，用来存储链表中每个节点的值
        Stack<Integer> stack = new Stack<>();
        //构建一个指针，指向链表的头结点位置，从它开始向后遍历
        ListNode curNode = head;
        //不断的遍历原链表中的每个节点，直到为null
        while ( curNode != null) {
            //把每个节点的值加入到栈中
            stack.push(curNode.val);
            //curNode向后移动
            curNode = curNode.next;
        }
        //获取栈的长度
        int size = stack.size();
        //定义一个同样长度的数组 res
        int[] res = new int[size];
        //遍历栈，从栈顶挨个弹出每个值，把这些值依次加入到数组res中
        for (int i = 0;i < size;i++) {
            //数组接收栈顶元素值
            res[i] = stack.pop();
        }
        //返回结果
        return res;
    }
}