// https://leetcode.com/submissions/detail/81974486/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
  var head, temp

  // if l1=null, then return l2. vice versa
  if (!l1 || !l2) 
    return l1 || l2

  // if l1.val > l2.val, switch l1 and l2
  if (l1.val > l2.val) 
    [l1, l2] = [l2, l1]

  head = l1
  while (l2) {
    if (!l1.next || l2.val < l1.next.val) {
      temp = l2.next
      l1 = insertNode(l1, l2)
      l2 = temp
    } else {
      if (l1.next)
        l1 = l1.next
    }
  }
  function insertNode(node1, node2) {
    node2.next = node1.next
    node1.next = node2
    return node2
  }
  return head
};


// 1 - 3 - 5 - 7
// 4 - 5 - 7 - 9
// 
// 1 - 3 - 4 - 5 -7
// 5 - 7 - 9
// 
// [0] l1
// [1,3,4] l2
// 
// 