package queuesturct

import "fmt"

type queue struct {
	Arr []interface{}
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) Push(_v ...interface{}) {
	q.Arr = append(q.Arr, _v...)
}

func (q *queue) Pop() (interface{}, error) {
	if len(q.Arr) == 0 {
		return nil, fmt.Errorf("none")
	}
	v := q.Arr[0]
	q.Arr = q.Arr[1:]
	return v, nil
}

func (q *queue) Length() int64 {
	return int64(len(q.Arr))
}
