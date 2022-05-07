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
    @return: a middle node of the linked list
    """
    def middleNode(self, head):
        if head is None:
            return None
        slow = head
        fast = slow.next
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next

        return slow