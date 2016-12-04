// https://leetcode.com/problems/remove-linked-list-elements/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 *
 * Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
 * Return: 1 --> 2 --> 3 --> 4 --> 5
 */
var removeElements = function(head, val) {
  while (head && head.val === val) {
    head = head.next
  }
  var node = head
  while (node && node.next) {
    while (node.next && node.next.val == val) {
      node.next = node.next.next
    }
    node = node.next
  }
  return head
};