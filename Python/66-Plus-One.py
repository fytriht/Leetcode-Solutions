class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        carry = 1
        for i in reversed(range(len(digits))):
            d = digits[i] + carry
            digits[i] = d % 10
            carry = 1 if d >= 10 else 0
        return digits if carry == 0 else [carry] + digits
