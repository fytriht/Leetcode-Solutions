class Solution {
public:

  void moveZeroes(vector<int>& nums) {

    for (int i = 0, j = 0, size = nums.size(); i < size; i++) {

      while (nums[i] == 0 && i < size - 1) i++;

      while (nums[j] != 0 && j < i) j++;

      swap(nums[i], nums[j]);
      
    }
  }

};