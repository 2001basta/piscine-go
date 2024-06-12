package piscine

func ListLast(l *List) interface{} {
	a := l.Head
	if a != nil {
		for a.Next != nil {
			a = a.Next
		}
		return a.Data
	} else {
		return nil
	}
}
