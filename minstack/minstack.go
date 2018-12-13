package minstack

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

type Stack struct {
	num int
	min int
}

type MinStack struct {
	Items []Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Items: make([]Stack, 0),
	}
}

func (this *MinStack) Push(x int) {
	if len(this.Items) > 0 {
		preMin := this.Items[len(this.Items)-1].min
		if x < preMin {
			this.Items = append(this.Items, Stack{
				num: x,
				min: x,
			})
		} else {
			this.Items = append(this.Items, Stack{
				num: x,
				min: preMin,
			})
		}
	} else {
		this.Items = append(this.Items, Stack{
			num: x,
			min: x,
		})
	}

}

func (this *MinStack) Pop() {
	this.Items = this.Items[0 : len(this.Items)-1]
}

func (this *MinStack) Top() int {
	return this.Items[len(this.Items)-1].num
}

func (this *MinStack) GetMin() int {
	return this.Items[len(this.Items)-1].min
}
