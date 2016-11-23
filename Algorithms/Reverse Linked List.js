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
 * Given head = [1,2,3,4,5], return [5,4,3,2,1]
 */
var reverseList = function(head) {
  var temp0, temp1
  if (head == null || head.next == null) {
    return head
  }
  if (head.next.next == null) {
    temp0 = head.next
    head.next.next = head
    head.next = null
    return temp0
  }
  function recur(node0, node1) { // null, 1 / 2, 3 / 4, 5
    if (node1.next == null) {
      node1.next = node0 // 5 -> 4
      return node1
    }
    temp0 = node1.next        // 2 / 4
    temp1 = node1.next.next  // 3 / 5(null)
    node1.next.next = node1   // 2 -> 1 / 4 -> 3
    node1.next = node0        // 1 -> null / 3 -> 2
    if (temp0 == null || temp1 == null) {
      return temp0
    }
    return recur(temp0, temp1) // 2, 3 / 4, 5
  }
  return recur(null, head)
};

// [1]
// [1, 2]
// [1, 2, 3]
// [1, 2, 3, 4]

