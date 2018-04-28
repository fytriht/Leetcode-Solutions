package solution

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
