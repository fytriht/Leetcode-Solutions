class Solution(object):
    def removeNthFromEnd(self, head, n):
        slow = head
        fast = head.next
        while n > 0:
            if not fast:
                return head.next
            fast = fast.next
            n -= 1
        while fast:
            fast = fast.next
            slow = slow.next
        slow.next = slow.next.next
        return head
