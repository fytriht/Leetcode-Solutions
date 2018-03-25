class Solution:
    def twoSum(self, nums, target):
        hash = {}
        for i, num in enumerate(nums):
            remind = target - num
            if remind in hash:
                return [hash[remind], i]
            hash[num] = i
        return []
