package piscine

func ListReverse(l *List) {
	c := l.Head
	if c != nil {
		l.Tail = l.Head
		var x *NodeL
		var y *NodeL
		for c != nil {
			x = c.Next
			c.Next = y
			y = c
			c = x
		}

		l.Head = y
	} else {
		return
	}
}
