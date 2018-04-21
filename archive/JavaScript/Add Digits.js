// https://leetcode.com/problems/add-digits/

/**
 * @param {number} num
 * @return {number}
 *
 * if num = 38, so 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.
 */
var addDigits = function(num) {
    return num && ( num % 9 || 9 )
}