let getIntersectionNode = (headA, headB) => {
  let currA = headA
  let currB = headB
  let tailA = (tailB = null)
  while (currA && currB) {
    if (currA === currB) {
      return currA
    }
    if (currA.next) {
      currA = currA.next
    } else if (!tailA) {
      tailA = currA
      currA = headB
    } else {
      break
    }
    if (currB.next) {
      currB = currB.next
    } else if (!tailB) {
      tailB = currB
      currB = headB
    } else {
      break
    }
  }
  return null
}
