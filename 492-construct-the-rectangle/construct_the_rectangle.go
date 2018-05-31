package solution

// TODO: gofmt
// TODO: test

func constructRectangle(area int) []int {
    s := int(math.Sqrt(float64(area)))
    ret := []int{s, s}
    for {
        if a := ret[0] * ret[1]; a < area {
            ret[0]++
        } else if a > area {
            ret[1]--
        } else {
            break
        }
    }
    return ret
}