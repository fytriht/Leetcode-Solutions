let sortString = s => [...s].sort().join()

let isAnagram = (s, t) => sortString(s) == sortString(t)
