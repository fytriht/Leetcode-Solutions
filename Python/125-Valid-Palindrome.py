class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        i, j = 0, len(s) - 1
        while i < j:
            while not s[i].isalnum():
                i += 1
                if i >= j:
                    return True
            while not s[j].isalnum():
                j -= 1
                if i >= j:
                    return True
            if s[i].lower() != s[j].lower():
                return False
            i, j = i + 1, j - 1
        return True
