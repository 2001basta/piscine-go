package piscine

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l.Head != nil {
		c := l.Head
		for c.Next != nil {
			if cond(c) {
				f(c)
			}
			c = c.Next
		}
		if cond(c) {
			f(c)
		}
	} else {
		return
	}
}
