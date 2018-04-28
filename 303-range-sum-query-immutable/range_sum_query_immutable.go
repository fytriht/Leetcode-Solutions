package solution

//
// solution 1
//

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums,
	}
}

func (n *NumArray) SumRange(i int, j int) int {
	var ret int
	for i <= j {
		ret += n.data[i]
		i++
	}
	return ret
}

//
// solution 2
//

type NumArray2 struct {
	sums []int
}

func Constructor2(nums []int) NumArray2 {
	s := make([]int, len(nums)+1)
	for i, n := range nums {
		s[i+1] = s[i] + n
	}
	return NumArray2{
		s,
	}
}

func (n *NumArray2) SumRange(i int, j int) int {
	return n.sums[j+1] - n.sums[i]
}
