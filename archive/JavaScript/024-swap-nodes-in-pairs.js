let swapPairs = head => {
    if (!head || !head.next) {
        return head
    }
    let tmp = head
    head = head.next
    tmp.next = swapPairs(tmp.next.next)
    head.next = tmp
    return head
}
