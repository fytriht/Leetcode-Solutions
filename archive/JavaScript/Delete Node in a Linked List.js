// https://leetcode.com/problems/delete-node-in-a-linked-list/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} node
 * @return {void} Do not return anything, modify node in-place instead.
 */
var deleteNode = function(node) {
  if (node.next == null) {        *// 
    node = null                   *// Without this part, the solution is also accepted by Leetcode.
  }                               *// I suppose this is a bug?
  while (node.next) {
    node.val = node.next.val
    if (node.next.next == null) {
      node.next = null
      break
    }
    node = node.next
  }
};


// 4 -> 5           5
// 5-> 6 -> 6      5 -> 6