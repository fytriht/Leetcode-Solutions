class Solution {
public:

  void moveZeroes(vector<int>& nums) {
    int i = 0, j = 0, len = nums.size();

    while (i < len) {

      while (nums[i] == 0 && i < len - 1) {
        i++;
      }

      while (nums[j] != 0 && j < i) {
        j++;
      }

      swap(nums[i], nums[j]);
      i++;
    }

  }
};