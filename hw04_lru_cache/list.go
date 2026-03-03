package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}, key string) *ListItem
	PushBack(v interface{}, key string) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Key   string
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Head *ListItem
	Tail *ListItem
	size int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}, key string) *ListItem {
	element := &ListItem{
		Value: v,
		Key:   key,
		Next:  l.Head,
		Prev:  nil,
	}
	if l.Head != nil {
		l.Head.Prev = element
	}
	l.Head = element

	if l.Tail == nil {
		l.Tail = element
	}

	l.size++
	return element
}

func (l *list) PushBack(v interface{}, key string) *ListItem {
	element := &ListItem{
		Value: v,
		Key:   key,
		Next:  nil,
		Prev:  l.Tail,
	}
	if l.Tail != nil {
		l.Tail.Next = element
	}
	l.Tail = element

	if l.Head == nil {
		l.Head = element
	}

	l.size++
	return element
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	prev := i.Prev
	next := i.Next

	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}

	if l.Head == i {
		l.Head = next
	}

	if l.Tail == i {
		l.Tail = prev
	}

	i.Prev = nil
	i.Next = nil

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value, i.Key)
}
