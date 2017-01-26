using namespace std;

class Solution {

public:
  int singleNumber(vector<int>& nums) {

    int result = 0;
    
    for (auto const &item : nums) result ^= item;

    return result;
  }

};