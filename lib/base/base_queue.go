package base

import (
	"container/list"
)

// Queue :
type Queue struct {
	list *list.List
}

// NewQueue :
func NewQueue() *Queue {
	qs := new(Queue)
	qs.list = list.New()
	return qs
}

// PushFront :
func (qs *Queue) PushFront(x interface{}) {
	qs.list.PushFront(x)
}

// PushBack :
func (qs *Queue) PushBack(x interface{}) {
	qs.list.PushBack(x)
}

// PopFront :
func (qs *Queue) PopFront() interface{} {
	x := qs.list.Front()
	return qs.list.Remove(x)
}

// PopBack :
func (qs *Queue) PopBack() interface{} {
	x := qs.list.Back()
	return qs.list.Remove(x)
}
