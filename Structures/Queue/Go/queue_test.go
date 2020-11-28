package queue

import "testing"

func TestQueue(t *testing.T) {
	tests := []struct {
		capacity int
		enqueues []int
		dequeues []int
	}{
		{
			5,
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		queue := NewQueue(test.capacity)
		for _, keys := range test.enqueues {
			err := queue.Enqueue(keys)
			if err != nil {
				t.Error(err)
			}
		}
		for _, key := range test.dequeues {
			dequeued, err := queue.Dequeue()
			if err != nil {
				t.Error(err)
			}
			if dequeued != key {
				t.Errorf("%v != %v ", dequeued, key)
			}
		}
	}
}

func TestOverflow(t *testing.T) {
	tests := []struct {
		capacity   int
		enqueues   []int
		overflowed bool
	}{
		{
			5,
			[]int{1, 2, 3, 4, 5},
			false,
		},
		{
			4,
			[]int{1, 2, 3, 4, 5},
			true,
		},
	}

	for _, test := range tests {
		queue := NewQueue(test.capacity)
		var err error
		for _, key := range test.enqueues {
			err = queue.Enqueue(key)
		}
		if (err != nil && !test.overflowed) || (err == nil && test.overflowed) {
			t.Error("there is a bug")
		}
	}
}

func TestUnderflow(t *testing.T) {
	queue := NewQueue(5)
	err := queue.Enqueue(4)
	if err == nil {
		t.Error("there is a bug in underflow")
	}
}
