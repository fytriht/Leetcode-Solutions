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
  var hash = {}, nodeA = headA, nodeB = headB

  while (nodeA || nodeB) {
    if (nodeA) {
      if (hash[nodeA.val] != undefined) {
        return nodeA
      }
      hash[nodeA.val] = nodeA.val
      nodeA = nodeA.next
    }
    if (nodeB) {
      if (hash[nodeB.val] != undefined) {
        return nodeB
      }
      hash[nodeB.val] = nodeB.val
      nodeB = nodeB.next
    }
  }
  return null
};