class Solution {
public:

  int addDigits(int num) {
    
    int sum = 0;
    long b = 10;

    while (num != 0) {
      int c = num % b;
      sum += c / (b / 10);
      num -= c;
      b *= 10;
    }
    
    return sum >= 10 ? addDigits(sum) : sum;
  }
  
};