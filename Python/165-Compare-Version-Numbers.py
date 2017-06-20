class Solution(object):
    def compareVersion(self, version1, version2):
        """
        :type version1: str
        :type version2: str
        :rtype: int
        """
        version1 = version1.split('.')
        version2 = version2.split('.')
        len1 = len(version1)
        len2 = len(version2)

        for i in range(max(len1, len2)):
            if i > len1 - 1:
                if int(version2[i]) != 0:
                    return -1
                continue
            if i > len2 - 1:
                if int(version1[i]) != 0:
                    return 1
                continue

            if int(version1[i]) != int(version2[i]):
                return 1 if int(version1[i]) > int(version2[i]) else -1
        return 0
