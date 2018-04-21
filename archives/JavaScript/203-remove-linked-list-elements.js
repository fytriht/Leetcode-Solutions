let removeElements = (head, val) =>
  head
    ? ((head.next = removeElements(head.next, val)),
      head.val == val ? head.next : head)
    : head
