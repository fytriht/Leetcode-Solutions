class Solution {
public:
  bool canConstruct(string ransomNote, string magazine) {
    vector<int> h(26);

    for (const auto& c : magazine) h[c - 'a']++;

    for (const auto& c : ransomNote) {
      if (--h[c - 'a'] < 0) return false;
    }

    return true;
  }
};