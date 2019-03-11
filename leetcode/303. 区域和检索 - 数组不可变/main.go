package main

type NumArray struct {
	Nums  []int
	LN    int
	Cache []int // 装的是0 - 的累加和的值
}

func Constructor(nums []int) NumArray {
	ln := len(nums)
	cache := make([]int, ln)
	for i, _ := range cache {
		cache[i] = -1
	}
	return NumArray{
		Nums:  nums,
		LN:    ln,
		Cache: cache,
	}
}

func (this *NumArray) Get(i int) int { // 0+1+2+...+i-1+i的值
	result := 0
	for k := i; k >= 0; k-- {
		if this.Cache[k] != -1 {
			result += this.Cache[k]
			break
		} else {
			result += this.Nums[k]
		}
	}
	if i >= 0 {
		this.Cache[i] = result
	}
	return result
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Get(j) - this.Get(i-1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {

}
