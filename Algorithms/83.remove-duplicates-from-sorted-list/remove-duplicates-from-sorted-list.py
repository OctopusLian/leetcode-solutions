# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        curt = head
        while curt:
            while curt.next and curt.next.val == curt.val:
                curt.next = curt.next.next
            curt = curt.next
        return head
