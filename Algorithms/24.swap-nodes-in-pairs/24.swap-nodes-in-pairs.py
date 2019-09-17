#
# @lc app=leetcode id=24 lang=python3
#
# [24] Swap Nodes in Pairs
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        pre,pre.next = self,head
        while pre.next and pre.next.next:
            a = pre.next
            b = pre.next
            pre.next,b.next,a.next = b,a,b.next
            pre = a

        return self.next

