package problem0622

type MyCircularQueue struct {
	Queue []int
	Size  int
}

/** Constructor initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Queue: make([]int, 0, k),
		Size:  k,
	}
}

/** EnQueue insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() == true {
		return false
	}

	this.Queue = append(this.Queue, value)

	return true
}

/** DeQueue delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() == true {
		return false
	}

	this.Queue = this.Queue[1:]

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() == true {
		return -1
	}

	return this.Queue[0]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() == true {
		return -1
	}

	return this.Queue[len(this.Queue)-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.Queue) == 0 {
		return true
	}

	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if len(this.Queue) == this.Size {
		return true
	}

	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
