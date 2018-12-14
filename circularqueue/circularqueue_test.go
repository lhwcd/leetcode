package circularqueue

import (
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestCircularQueue(t *testing.T) {

	fmt.Println((-5) % 4)
	fmt.Println((0 + 3) % 4)

	circularQueue := Constructor(3)
	assert.Equal(t, true, circularQueue.EnQueue(1))
	assert.Equal(t, true, circularQueue.EnQueue(2))
	assert.Equal(t, true, circularQueue.EnQueue(3))
	assert.Equal(t, false, circularQueue.EnQueue(4))
	assert.Equal(t, 3, circularQueue.Rear())
	assert.Equal(t, true, circularQueue.IsFull())
	assert.Equal(t, true, circularQueue.DeQueue())
	assert.Equal(t, true, circularQueue.EnQueue(4))
	assert.Equal(t, 4, circularQueue.Rear())

}
