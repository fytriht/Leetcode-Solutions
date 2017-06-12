# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = ListNode(0)
        curr = head

        while l1 or l2:
            if not l1:
                while l2:
                    curr.next = l2
                    curr = curr.next
                    l2 = l2.next
                return head.next
            if not l2:
                while l1:
                    curr.next = l1
                    curr = curr.next
                    l1 = l1.next
                return head.next
            if l1.val <= l2.val:
                curr.next = l1
                l1 = l1.next
            else:
                curr.next = l2
                l2 = l2.next
            curr = curr.next
        return head.next