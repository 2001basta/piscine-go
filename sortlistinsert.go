package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}

	if l != nil {
		n.Next = l
		n = ListSort(n)
	} else {
		l = n
		return n
	}
	return n
}
