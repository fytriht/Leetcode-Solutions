let getIntersectionNode = (headA, headB) => {
  let currA = headA
  let currB = headB
  let aReached = (bReached = false)
  while (currA && currB) {
    if (currA === currB) {
      return currA
    }
    currA = currA.next
    currB = currB.next
    if (!currA && !aReached) {
      aReached = true
      currA = headB
    }
    if (!currB && !bReached) {
      bReached = true
      currB = headA
    }
  }
  return null
}
