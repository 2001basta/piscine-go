package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	a := l
	count := 0

	if a != nil {
		if pos == 0 {
			return l
		} else {
			for a.Next != nil {
				count++
				a = a.Next
				if count == pos {
					return a
				}
			}
		}
	} else {
		return nil
	}
	return nil
}
