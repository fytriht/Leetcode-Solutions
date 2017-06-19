class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        result = []
        for i in range(numRows):
            result.append(self.genRow(i + 1))
        return result

    def genRow(self, row):
        if row == 1:
            return [1]
        if row == 2:
            return [1, 1]
        last = self.genRow(row -1)
        i = 1
        result = []
        while i < row - 1:
            result.append(last[i] + last[i - 1])
            i += 1
        return [1] + result + [1]