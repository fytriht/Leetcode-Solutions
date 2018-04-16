let reverseList = head => {
  if (!head || !head.next) return head
  let r = reverseList(head.next)
  head.next.next = head
  head.next = null
  return r
}
