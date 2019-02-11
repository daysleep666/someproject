package main

import (
	"fmt"
	"time"
)

type choiceType string

type Task struct {
	Index    int64
	Task     func()
	RunSign  chan string
	StopSign chan string
}

func NewTask(_index int64) *Task {
	task := &Task{
		Index:    _index,
		RunSign:  make(chan string),
		StopSign: make(chan string),
	}
	go task.Run()
	return task
}

func (task *Task) Run() {
	for {
		// 监听chan的消息
		select {
		case <-task.RunSign:
			if task.Task != nil {
				task.Task()
				task.Task = nil
			}

		case <-task.StopSign:
			return
		}
	}
}

func (task *Task) Start(_task func()) {
	task.Task = _task
	task.RunSign <- "start"
}

func (task *Task) Stop() {
	task.StopSign <- "stop"
}

// 队列
type Queue struct {
	Task []*Task
}

func NewQueue(_count int64) *Queue {
	var queue Queue
	for i := int64(0); i < _count; i++ {
		queue.push(NewTask(i))
	}
	return &queue
}

func (q *Queue) push(_task *Task) {
	q.Task = append(q.Task, _task)
}

func (q *Queue) pop() (*Task, error) {
	if len(q.Task) == 0 {
		return nil, fmt.Errorf("队列已空 ")
	}
	task := q.Task[0]
	q.Task = q.Task[1:]
	return task, nil
}

func (q *Queue) Insert(_task func()) error {
	task, err := q.pop()
	if err != nil {
		return err
	}
	task.Start(_task)
	return nil
}

func main() {
	queue := NewQueue(3)
	queue.Insert(func() {
		fmt.Println("啦啦啦")
		time.Sleep(time.Second * 3)
	})
	queue.Insert(func() {
		fmt.Println("啦啦啦")
		time.Sleep(time.Second * 3)
	})
	queue.Insert(func() {
		fmt.Println("啦啦啦")
		time.Sleep(time.Second * 3)
	})
	queue.Insert(func() {
		fmt.Println("啦啦啦")
		time.Sleep(time.Second * 3)
	})

	time.Sleep(5 * time.Second)
}
