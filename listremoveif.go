package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	c := l.Head
	var d *NodeL
	for c != nil {
		if c.Data == data_ref {
			if d == nil {
				l.Head = c.Next
			} else {
				d.Next = c.Next
			}
		} else {
			d = c
		}
		c = c.Next
	}
}
