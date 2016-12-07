// https://leetcode.com/problems/intersection-of-two-linked-lists/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
  var hash = {}

  while (headA || headB) {
    if (headA) {
      if (hash[headA.val] != undefined) {
        return headA
      }
      hash[headA.val] = headA.val
      headA = headA.next
    }
    if (headB) {
      if (hash[headB.val] != undefined) {
        return headB
      }
      hash[headB.val] = headB.val
      headB = headB.next
    }
  }
  return null
};