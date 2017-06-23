class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cache = {}
        length = len(nums)
        for num in nums:
            if num in cache:
                cache[num] += 1
            else:
                cache[num] = 1
            if cache[num] > length / 2:
                return num
