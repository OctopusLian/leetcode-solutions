#
# @lc app=leetcode id=141 lang=python3
#
# [141] Linked List Cycle
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        fast = slow = head
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next
            if slow is fast:
                return True

        return False

