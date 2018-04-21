class Solution(object):
    def lengthOfLastWord(self, s):
        length = 0
        for ch in reversed(s):
            if ch == ' ':
                if length:
                    break
            else:
                length += 1
        return length
