class Solution(object):
    def rotate(self, nums, k):
        length = len(nums)
        k %= length
        last = nums[-k:]
        for i in reversed(range(length - k)):
            nums[i+k] = nums[i]
        for i in range(len(last)):
            nums[i] = last[i]
