class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        return x >= 0 and x == self.reverse(x)

    def reverse(self, x):
        result = 0
        while (x > 0):
            result = result * 10 + x % 10
            x //= 10
        return result
