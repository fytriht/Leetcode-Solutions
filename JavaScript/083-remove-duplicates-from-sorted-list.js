let deleteDuplicates = head => {
  for (let i = node; i; i = i.next) {
    while (i.next && i.val === i.next.val) {
      i.next = i.next.next
    }
  }
  return head
}
