class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ""

        firstStr = strs[0]
        for i in range(len(firstStr)):
            for string in strs[1:]:
                if len(string) <= i:
                    return string
                if string[i] != firstStr[i]:
                    return string[:i]
        return firstStr
