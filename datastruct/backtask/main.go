package main

import (
	"fmt"
	"time"
)

type OneNode struct {
	Key         string
	Task        func()
	ExecuteTime int64
	Next        *OneNode
}

func AddNode(_oneNode *OneNode, _key string, _task func(), _executeTime int64) *OneNode {
	if _oneNode == nil {
		newNode := &OneNode{Key: _key, Task: _task, ExecuteTime: _executeTime}
		return newNode
	}

	if _oneNode.ExecuteTime > _executeTime {
		newNode := &OneNode{Key: _key, Task: _task, ExecuteTime: _executeTime}
		return newNode
	}

	_oneNode.Next = AddNode(_oneNode.Next, _key, _task, _executeTime)
	return _oneNode
}

func Pop(_oneNode *OneNode) *OneNode {
	if _oneNode == nil {
		return nil
	}
	return _oneNode.Next
}

type BackTask struct {
	Head *OneNode
}

func (bt *BackTask) Insert(_key string, _task func(), _after int64) {
	bt.Head = AddNode(bt.Head, _key, _task, time.Now().Add(time.Second*time.Duration(_after)).Unix())
}

func (bt *BackTask) Run() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C
		func() {
			if bt.Head != nil {
				if bt.Head.ExecuteTime < time.Now().Unix() {
					go bt.Head.Task()
					bt.Head = Pop(bt.Head)
				}
			}
		}()
	}
}

func main() {
	var bt = new(BackTask)
	bt.Insert("a", func() { fmt.Println("tasl1", time.Now().Unix()) }, 1)
	bt.Insert("a", func() { fmt.Println("tasl2", time.Now().Unix()) }, 4)
	bt.Insert("a", func() { fmt.Println("tasl3", time.Now().Unix()) }, 10)
	bt.Run()
}
