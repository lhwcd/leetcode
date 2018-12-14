package numberofrecentcalls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	minVal := t - 3000
	this.queue = append(this.queue, t)
	//
	for k, v := range this.queue {
		if v >= minVal {
			this.queue = this.queue[k:len(this.queue)]
			break
		}
	}
	//
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
