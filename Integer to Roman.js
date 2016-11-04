// https://leetcode.com/problems/integer-to-roman/

/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function(num) {
  
  const pairs = [['1', 'I'], ['5', 'V'], ['10', 'X'],
                  ['50', 'L'], ['100', 'C'], ['500', 'D'], ['1000', 'M']]
  
  const map = new Map(pairs)

  const romanizer = (item, idx) => {

    item = String(item)

    if (map.has(item))
      return map.get(item)
    
    if (item[0] == '9')
      return romanizer(Math.pow(10, idx)) + 
               romanizer(Number(item) + Math.pow(10, idx))
  
    if (item[0] <= '3')
      return romanizer(Math.pow(10, idx)).repeat(item[0])
    
    if (item[0] == '4')
      return romanizer(Math.pow(10, idx)) + 
               romanizer(Number(item) + Math.pow(10, idx))
    
    return romanizer(item.replace(/\d/, '5')) + 
             romanizer(Math.pow(10, idx)).repeat(item[0] - '5')

  }

    // 1950 => [0, 50, 900, 1000]
  return String(num).split('').reverse()
           .map((item, idx) => item * Math.pow(10, idx))
           .map(romanizer).reverse().join('')  
};