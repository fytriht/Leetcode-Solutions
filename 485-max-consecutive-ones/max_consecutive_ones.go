package solution

// TODO: gofmt
// TODO: test

func findMaxConsecutiveOnes(nums []int) int {
    var ret, curr int
    var match bool
    for _, n := range nums {
        if n == 1 {
            match = true
            curr++
        } else if match {
            ret = max(ret, curr)
            match = false
            curr = 0
        }
    }
    return max(curr, ret)
}

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}