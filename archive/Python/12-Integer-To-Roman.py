class Solution(object):
    def intToRoman(self, num):
        map = {
            1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC",
            50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"
        }
        result = ''
        keys = sorted(map.keys())
        for key in reversed(keys):
            q = num // key
            num -= q * key
            if q > 0:
                result += map[key] * q
        return result
