# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        nodeA, nodeB = headA, headB
        cache = {}

        while nodeA or nodeB:
            if nodeA:
                if nodeA in cache:
                    return nodeA
                else:
                    cache[nodeA] = nodeA
                    nodeA = nodeA.next
            if nodeB:
                if nodeB in cache:
                    return nodeB
                else:
                    cache[nodeB] = nodeB
                    nodeB = nodeB.next
        return None
