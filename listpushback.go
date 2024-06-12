package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head != nil {
		a := l.Head
		for a.Next != nil {
			a = a.Next
		}
		a.Next = n
	} else {
		l.Head = n
		return
	}
}
