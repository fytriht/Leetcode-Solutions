class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        length = 0
        for i, ch in enumerate(s):
            if ch != ' ':
                if s[i-1] == ' ':
                    length = 0
                length += 1
        return length
