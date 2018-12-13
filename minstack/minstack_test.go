package minstack

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

// * obj := Constructor();
// * obj.Push(x);
// * obj.Pop();
// * param_3 := obj.Top();
// * param_4 := obj.GetMin();

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.GetMin())
	minStack.Pop()
	assert.Equal(t, 0, minStack.Top())
	assert.Equal(t, -2, minStack.GetMin())
}
