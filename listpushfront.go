package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head != nil {
		n.Next = l.Head
		l.Head = n
	} else {
		l.Head = n
	}
}
