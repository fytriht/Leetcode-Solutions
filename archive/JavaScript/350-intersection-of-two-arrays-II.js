let intersect = (nums1, nums2) => {
  if (nusm1.length > nums2.length) {
    return intersect(nums2, nums1)
  }
  let m = {}
  let ret = []
  nums1.forEach(n => {
    if (m[n] == undefined) m[n] = 0
    m[n]++
  })
  nums2.forEach(n => {
    if (m[n] != undefined && m[n] != 0) {
      ret.push(n)
      m[n]--
    }
  })
  return ret
}
