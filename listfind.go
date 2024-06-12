package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	c := l.Head
	if c != nil {
		if comp(c.Data, ref) {
			return &c.Data
		} else {
			for c.Next != nil {
				c = c.Next
				if comp(c.Data, ref) {
					return &c.Data
				}
			}
		}
	}
	return nil
}
