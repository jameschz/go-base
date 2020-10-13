package base

import "container/list"

// Stack :
type Stack struct {
	list *list.List
}

// NewStack :
func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

// Push :
func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

// Pop :
func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		return stack.list.Remove(e)
	}
	return nil
}

// Peak :
func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

// Len :
func (stack *Stack) Len() int {
	return stack.list.Len()
}

// Empty :
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}
