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
 *
 * Given 1->2->3->4->5, n = 2, return 1->2->3->5
 */
var removeNthFromEnd = function(head, n) {
  var arr = [], len
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