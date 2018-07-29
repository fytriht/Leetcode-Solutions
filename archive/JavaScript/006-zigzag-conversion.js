let convert = (s, numRows) => {
    if (numRows == 1) {
        return s
    }
    let ret = ""
    for (let i = 0; i < numRows; i++) {
        for (let j = i; j < s.length; j += 2 * numRows - 2) {
            ret += s[j]
            let k = 2 * numRows - 2 + j - 2 * i
            if (k < s.length && i != 0 && i != numRows - 1) {
                ret += s[k]
            }
        }
    }
    return ret
}
