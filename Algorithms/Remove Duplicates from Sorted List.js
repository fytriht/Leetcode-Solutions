// https://leetcode.com/problems/remove-duplicates-from-sorted-list/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 *
 * Given 1->1->1->2->3->3, return 1->2->3.
 */
var deleteDuplicates = function(head) {
  if (!head) return []
  var node = head
  while (node && node.next) {
    while(node.next && node.next.val == node.val) {
      node.next = node.next.next
    }
    node = node.next
  }
  return head
};
