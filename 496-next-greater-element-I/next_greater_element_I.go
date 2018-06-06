package solution

// TODO: test

//
// solution 1
//

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

//
// solution 2
//

func nextGreaterElement2(findNums []int, nums []int) []int {
	m := make(map[int]int)
	stack := []int{}
	for _, n := range nums {
		for len(stack) > 0 {
			if top := stack[len(stack)-1]; top < n {
				m[top] = n
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, n)
	}

	var ret []int
	for _, n := range findNums {
		if v, ok := m[n]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}
