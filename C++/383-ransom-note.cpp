class Solution {
public:
  bool canConstruct(string ransomNote, string magazine) {
    vector<int> h(26);

    for (const auto& c : magazine) {
      h[c - 'a']++;
    }

    for (const auto& c : ransomNote) {
      h[c - 'a']--;
    }

    for (const auto& c : h) {
      if (c < 0) return false;
    }
    
    return true;
  }
};