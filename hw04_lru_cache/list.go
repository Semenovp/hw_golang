package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	itemsMap map[*ListItem]*ListItem
	len      int
	front    *ListItem
	back     *ListItem
}

func NewList() List {
	newList := new(list)
	newList.itemsMap = make(map[*ListItem]*ListItem)
	return newList
}

func (l *list) firstElem(v interface{}) *ListItem {
	firstElem := ListItem{v, nil, nil}
	l.front = &firstElem
	l.back = &firstElem
	l.len = 1
	return &firstElem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{
		Value: v,
		Next:  l.Front(),
		Prev:  nil,
	}

	if l.len == 0 {
		l.back = newListItem
	}
	if l.Front() != nil {
		l.Front().Prev = newListItem
	}
	l.itemsMap[newListItem] = newListItem
	l.front = newListItem
	l.len++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.len == 0 {
		return l.firstElem(v)
	}
	var newElem ListItem

	newElem.Value = v
	newElem.Prev = l.back
	l.back.Next = &newElem
	l.back = &newElem
	l.len++
	return &newElem
}

func (l *list) Remove(i *ListItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if l.len > 0 {
		l.len--
	}
	delete(l.itemsMap, i)
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Front() {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		if i.Next == nil {
			l.back = i.Prev
		}
		i.Prev.Next = i.Next
	}

	l.Front().Prev = i
	i.Next = l.Front()
	i.Prev = nil
	l.front = i
}
