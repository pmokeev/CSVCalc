package queue

import "testing"

func TestQueue(t *testing.T) {
	tests := []struct {
		name      string
		doFunc    func(q *Queue)
		checkFunc func(q *Queue) bool
	}{
		{
			name: "Success: simple test",
			doFunc: func(q *Queue) {
				q.Push(&Term{
					XKey: "1",
				})
			},
			checkFunc: func(q *Queue) bool {
				popped := q.Pop()

				return popped.XKey == "1"
			},
		},
		{
			name: "Success: more values in a queue",
			doFunc: func(q *Queue) {
				q.Push(&Term{
					XKey: "1",
				})
				q.Push(&Term{
					XKey: "2",
				})
				q.Push(&Term{
					XKey: "3",
				})
				q.Push(&Term{
					XKey: "4",
				})
			},
			checkFunc: func(q *Queue) bool {
				for _, value := range []string{"1", "2", "3", "4"} {
					popped := q.Pop()

					if popped.XKey != value {
						return false
					}
				}

				return q.Empty()
			},
		},
		{
			name: "Success: more complex test",
			doFunc: func(q *Queue) {
				q.Push(&Term{
					XKey: "1",
				})
				q.Push(&Term{
					XKey: "2",
				})
				q.Push(&Term{
					XKey: "3",
				})
				q.Pop()
				q.Pop()
				q.Push(&Term{
					XKey: "4",
				})
			},
			checkFunc: func(q *Queue) bool {
				for _, value := range []string{"3", "4"} {
					popped := q.Pop()

					if popped.XKey != value {
						return false
					}
				}

				return q.Empty()
			},
		},
		{
			name: "Success: empty queue after pushes",
			doFunc: func(q *Queue) {
				q.Push(&Term{
					XKey: "1",
				})
				q.Push(&Term{
					XKey: "2",
				})
				q.Push(&Term{
					XKey: "3",
				})
				q.Push(&Term{
					XKey: "4",
				})
				q.Pop()
				q.Pop()
				q.Pop()
				q.Pop()
			},
			checkFunc: func(q *Queue) bool {
				return q.Empty()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()

			tt.doFunc(q)

			if !tt.checkFunc(q) {
				t.Error("Push() = invalid value in queue")
			}
		})
	}
}
