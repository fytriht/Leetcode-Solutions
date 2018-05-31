package solution 

// TODO: gofmt
// TODO: test

func nextGreaterElement(findNums []int, nums []int) []int {
    var ret []int
    for i, n1 := range findNums {
        var match bool
        ret = append(ret, -1)
        for _, n2 := range nums {
            if match && n2 > n1 {
                ret[i] = n2
                break
            }
            if n1 == n2 {
                match = true
            }   
        }
    }
    return ret
}