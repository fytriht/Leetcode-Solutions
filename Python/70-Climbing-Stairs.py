class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0 or n == 1:
            return n
        a, b = 0, 1
        for i in range(n):
            a, b =  b, a + b
        return b
