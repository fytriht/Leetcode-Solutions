class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        length = 0
        for ch in reversed(s):
            if ch == ' ':
                if length:
                    break
            else:
                length += 1
        return length
