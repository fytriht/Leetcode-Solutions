class Solution(object):
    def romanToInt(self, s):
        map = {
            'I': 1, 'IV': 4, 'V': 5, 'IX': 9, 'X': 10, 'XL': 40, 'L': 50,
            'XC': 90, 'C': 100, 'CD': 400, 'D': 500, 'CM': 900, 'M': 1000
        }
        result = 0
        length = len(s)
        if length == 1:
            return map[s] or result
        for i in range(length):
            if s[i:i+2] in map and i < length -1:
                result += map[s[i:i+2]]
            elif s[i] in map and s[i-1:i+1] not in map:
                result += map[s[i]]
        return result
