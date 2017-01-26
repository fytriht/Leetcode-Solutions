#ifndef Solution_h
#define Solution_h

using namespace std;

class Solution {

public:
  string reverseString (string s) {
    
    string result;
    int len = s.size();

    for (int i = len - 1; i >= 0; i--) {
      result += s[i];
    }

    return result;
  }

};

#endif