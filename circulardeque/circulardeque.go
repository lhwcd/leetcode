package circulardeque

//MyCircularDeque four param
type MyCircularDeque struct {
	front int
	rear  int
	size  int
	items []int
}

//Constructor return a new MyCircularDeque with size k
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		front: 0,
		rear:  0,
		size:  k + 1,
		items: make([]int, k+1),
	}
}

//IsEmpty return true if empty
func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.rear+1)%this.size == this.front ||
		(this.front-1+this.size) == this.rear
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.front]
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + this.size) % this.size
	this.items[this.front] = value
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.size
	return true
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[(this.rear-1+this.size)%this.size]
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.items[this.rear] = value
	this.rear = (this.rear + 1) % this.size
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.rear = (this.rear - 1 + this.size) % this.size
	return true
}
