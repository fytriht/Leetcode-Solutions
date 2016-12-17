// https://leetcode.com/submissions/detail/85913288/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
  if (head === null || head.next === null) {
    return false
  }

  var 
    slow = head,
    fast = head

  while (fast.next && fast.next.next) {
    fast = fast.next.next
    slow = slow.next
    if (fast == slow) {
      return true
    }
  }

  return false
};