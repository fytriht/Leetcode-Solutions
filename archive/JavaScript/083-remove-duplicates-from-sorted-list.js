let deleteDuplicates = head => (
  head &&
    head.next &&
    (head.val === head.next.val
      ? (head = deleteDuplicates(head.next))
      : (head.next = deleteDuplicates(head.next))),
  head
)
