package solution

// TODO: gofmt
// TODO: test

func findMaxConsecutiveOnes(nums []int) int {
    var ret, curr int
    for _, n := range nums {
        if n == 1 {
            curr++
        } else {
            curr = 0
        }
        ret = max(ret, curr)
    }
    return ret
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}