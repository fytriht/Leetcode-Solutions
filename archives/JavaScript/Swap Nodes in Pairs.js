// https://leetcode.com/problems/swap-nodes-in-pairs/

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
 */
var swapPairs = function(head) {
  // if  head=null, returnnull
  if (!head) 
    return null
  // if list=[head] return head
  if (!head.next) 
    return head
  var node = head, temp, temp1  
  head = node.next
  while (node && node.next) {
    temp1 = null
    temp = null
    if (node.next.next) {
      temp = node.next.next //3
      temp1 = node.next.next.next || node.next.next //null
    }
    node.next.next = node //2->1
    node.next = temp1 //1->4
    node = temp //node=3
  }
  return head
};
// 1->2->3
// 2->1->3
// 
// 1->2->3->4 ->5->6
// 2->1->4->3 ->6->5
// 
// 2->1->4->3->6->5->null
// 
// 2->1->4
// 
// [1,2,3]
