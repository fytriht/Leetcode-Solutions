class Solution(object):
    def titleToNumber(self, s):
        """
        :type s: str
        :rtype: int
        """
        base = 1
        result = 0
        for i in reversed(range(len(s))):
            result += (ord(s[i]) - ord('A') + 1) * base
            base *= 26
        return result
