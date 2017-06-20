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
        times = max(len(version1), len(version2))

        for i in range(times):
            if i > len1 - 1:
                if int(version2[i]) != 0:
                    return -1
                continue
            if i > len2 - 1:
                if int(version1[i]) != 0:
                    return 1
                continue

            v1 = int(version1[i])
            v2 = int(version2[i])
            if v1 > v2:
                return 1
            if v2 > v1:
                return -1
        return 0
