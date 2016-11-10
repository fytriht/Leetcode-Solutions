// https://leetcode.com/submissions/detail/81974486/

/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {

 // if l1.val > l2.val, switch l1 and l2
  if (l1 && l2 && l1.val > l2.val) 
    [l1, l2] = [l2, l1]

  var head = l1, temp

  while (l2 && l1) {
    if (!l1.next || l2.val < l1.next.val) {
      l2 = insertNode(l1, l2)
    }
    l1 = l1.next || l1
  }
  function insertNode(node1, node2) {
    temp = node2.next
    node2.next = node1.next
    node1.next = node2
    return temp
  }
  return head || l2
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