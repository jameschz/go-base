package base

import "container/list"

// List :
type List struct {
	list *list.List
}

// NewList :
func NewList() *List {
	list := list.New()
	return &List{list}
}

// Back :
func (list *List) Back() *list.Element {
	return list.list.Back()
}

// BackValue :
func (list *List) BackValue() interface{} {
	return list.list.Back().Value
}

// Front :
func (list *List) Front() *list.Element {
	return list.list.Front()
}

// FrontValue :
func (list *List) FrontValue() interface{} {
	return list.list.Front().Value
}

// Len :
func (list *List) Len() int {
	return list.list.Len()
}

// InsertAfter :
func (list *List) InsertAfter(v interface{}, m *list.Element) *list.Element {
	return list.list.InsertAfter(v, m)
}

// InsertBefore :
func (list *List) InsertBefore(v interface{}, m *list.Element) *list.Element {
	return list.list.InsertBefore(v, m)
}

// MoveAfter :
func (list *List) MoveAfter(e *list.Element, m *list.Element) {
	list.list.MoveAfter(e, m)
}

// MoveBefore :
func (list *List) MoveBefore(e *list.Element, m *list.Element) {
	list.list.MoveBefore(e, m)
}

// MoveToBack :
func (list *List) MoveToBack(e *list.Element) {
	list.list.MoveToBack(e)
}

// MoveToFront :
func (list *List) MoveToFront(e *list.Element) {
	list.list.MoveToFront(e)
}

// PushBack :
func (list *List) PushBack(v interface{}) *list.Element {
	return list.list.PushBack(v)
}

// PushFront :
func (list *List) PushFront(v interface{}) *list.Element {
	return list.list.PushFront(v)
}

// PushBackList :
func (list *List) PushBackList(l *list.List) {
	list.list.PushBackList(l)
}

// PushFrontList :
func (list *List) PushFrontList(l *list.List) {
	list.list.PushFrontList(l)
}

// Remove :
func (list *List) Remove(e *list.Element) interface{} {
	return list.list.Remove(e)
}
