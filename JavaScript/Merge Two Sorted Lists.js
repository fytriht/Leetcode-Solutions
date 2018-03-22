let mergeTwoLists = (l1, l2) => {
  if (!l1 || !l2) return l1 || l2

  let head, curr
  if (l1.val < l2.val) {
    head = curr = l1
    l1 = l1.next
  } else {
    head = curr = l2
    l2 = l2.next
  }

  while (l1 && l2) {
    if (l1.val < l2.val) {
      curr.next = l1.next
      l1 = l1.next
    } else {
      curr.next = l2.next
      l2 = l2.next
    }
    curr = curr.next
  }

  curr.next = l1 || l2

  return head.next
}
