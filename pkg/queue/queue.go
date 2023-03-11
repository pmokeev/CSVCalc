package queue

type node struct {
	value    *Term
	previous *node
	next     *node
}

type Queue struct {
	head *node
	tail *node
	size int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
	}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

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

func (q *Queue) Pop() *Term {
	if q.head == nil {
		return nil
	}

	popedNode := q.head
	q.head = q.head.previous
	q.size--

	return popedNode.value
}
