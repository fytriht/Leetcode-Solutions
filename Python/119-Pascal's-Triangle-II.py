class Solution(object):
    def getRow(self, rowIndex):
        """
        :type rowIndex: int
        :rtype: List[int]
        """
        if rowIndex == 0:
            return [1]
        if rowIndex == 1:
            return [1, 1]

        result = [1]
        last = self.getRow(rowIndex - 1)
        for i in range(rowIndex - 1):
            result.append(last[i + 1] + last[i])
        result.append(1)
        return result
