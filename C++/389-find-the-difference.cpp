class Solution {
public:

  char findTheDifference(string s, string t) {
    
    char result = t[t.size() - 1];
    
    for (int i = 0, j = s.size() - 1; i <= j; i++, j--) {
      result ^= s[i] ^ t[i];
      if (i != j) result ^= s[j] ^ t[j];
    }

    return result;
  } 

};