package piscine

func ListMerge(l1 *List, l2 *List) {
	c := l2.Head
	a := l1.Head
	if a != nil {

		for a.Next != nil {
			a = a.Next
		}
		a.Next = c

	} else {
		l1.Head = c
		l1.Tail = l2.Tail
	}
}
