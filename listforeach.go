package piscine

func ListForEach(l *List, f func(*NodeL)) {
	if l.Head != nil {
		c := l.Head
		f(c)
		for c.Next != nil {
			c = c.Next
			f(c)
		}
	} else {
		return
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
