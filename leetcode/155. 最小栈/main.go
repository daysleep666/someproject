package main

type MinStack struct {
	Nums []int
	Min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Nums: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if this.Min > x {
		this.Min = x
	}
	if len(this.Nums) == 0 {
		this.Min = x
	}
	this.Nums = append(this.Nums, x)
}

func (this *MinStack) Pop() {
	if this.Top() == this.Min {
		this.Min = 0
		if len(this.Nums) > 0 {
			this.Min = this.Nums[0]
		}
		for i := 1; i < len(this.Nums)-1; i++ {
			if this.Nums[i] < this.Min {
				this.Min = this.Nums[i]
			}
		}
	}
	this.Nums = this.Nums[:len(this.Nums)-1]
}

func (this *MinStack) Top() int {
	return this.Nums[len(this.Nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
