// https://leetcode.com/submissions/detail/81325338/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
  let arr = [], len
  
  if (!head.next)
    return []
  
  while (head) {
    arr.push(head)
    head = head.next
  }
  
  len = arr.length
  if (n == len)
    return arr[1]
  
  arr[len - n - 1].next = arr[len - n].next
  
  return arr[0]
};