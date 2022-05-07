"""
Definition of ListNode
class ListNode(object):
    def __init__(self, val, next=None):
        self.val = val
        self.next = next
"""

class Solution:
    """
    @param head: the head of linked list.
    @return: An integer list
    """
    def toArrayList(self, head):
        # write your code here
        r = []
        #遍历链表，并将每个节点的val存储在数组中
        while head is not None:
            r.append(head.val)
            head = head.next
        return r