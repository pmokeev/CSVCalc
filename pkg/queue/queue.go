package queue

// node represents queue node.
type node struct {
	value    *Term
	previous *node
	next     *node
}

// Queue represents queue based on linked list.
type Queue struct {
	head *node
	tail *node
	size int
}

// NewQueue returns new instance of queue.
func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
	}
}

// Empty checks the queue for emptyness.
func (q *Queue) Empty() bool {
	return q.size == 0
}

// Size returns queue size.
func (q *Queue) Size() int {
	return q.size
}

// Push push value into queue.
func (q *Queue) Push(value *Term) {
	currentNode := &node{
		value: value,
	}
	q.size++

	if q.tail == nil {
		q.head = currentNode
		q.tail = currentNode
		return
	}

	currentNode.next = q.tail
	q.tail.previous = currentNode
	q.tail = currentNode
}

// Pop pops values from queue.
func (q *Queue) Pop() *Term {
	if q.head == nil {
		return nil
	}

	popedNode := q.head
	q.head = q.head.previous
	q.size--

	return popedNode.value
}
