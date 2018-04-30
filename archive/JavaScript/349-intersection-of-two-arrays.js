let intersection = (nums1, nums2) => {
  if (nums1.length > nums2.length) {
    return intersection(num2, nums1)
  }
  let s = new Set()
  let ret = []
  nums1.forEach(n => s.add(n))
  nums2.forEach(n => {
    if (s.has(n)) ret.push(n)
    s.delete(n)
  })
  return ret
}
