class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        carry = 0
        result = ''
        lengthA, lengthB = len(a), len(b)
        for i in range(max(lengthA, lengthB)):
            val = carry
            if i < lengthA:
                val += int(a[- (i + 1)])
            if i < lengthB:
                val += int(b[- (i + 1)])
            result += str(val % 2)
            carry = val // 2
        if carry == 1:
            result += '1'
        return result[::-1]