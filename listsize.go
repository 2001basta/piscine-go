package piscine

func ListSize(l *List) int {
	a := l.Head

	if a != nil {
		count := 1
		for a.Next != nil {
			a = a.Next
			count++
		}
		return count
	} else {
		return 0
	}
}
