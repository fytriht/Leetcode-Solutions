let removeNthFromEnd = (head, n) => {
    let slow = head, fast = head
    for (; n > 0; n--) fast = fast.next
    if (!fast) return head.next
    for (; fast.next; fast = fast.next, slow = slow.next) {
    }
    slow.next = slow.next.next
    return head
}
