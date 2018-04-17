let containsDuplicate = nums => nums.sort().some((n, i) => n == nums[i + 1])
