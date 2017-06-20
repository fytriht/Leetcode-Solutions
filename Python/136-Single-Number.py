class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cache = {}
        for num in nums:
            if num not in cache:
                cache[num] = 1
            else:
                cache[num] += 1
        for key, val in cache.items():
            if val == 1:
                return key
        return None
