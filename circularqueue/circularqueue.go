package circularqueue

type MyCircularQueue struct {
	items []int
	front int
	rear  int
	size  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		items: make([]int, k+1),
		front: 0,
		rear:  0,
		size:  k + 1,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	//
	if this.IsFull() {
		return false
	}
	//
	this.items[this.rear] = value
	this.rear = (this.rear + 1) % this.size
	return true
	//
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	//
	this.front = (this.front + 1) % this.size
	//
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	//
	if this.IsEmpty() {
		return -1
	}
	return this.items[(this.rear+this.size-1)%this.size]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if (this.rear)%this.size == this.front%this.size {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.rear+1)%this.size == this.front%this.size {
		return true
	}
	return false
}
