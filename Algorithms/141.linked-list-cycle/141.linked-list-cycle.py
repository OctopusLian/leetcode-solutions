class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        fast = slow = head
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next
            if slow is fast:
                return True

        return False

