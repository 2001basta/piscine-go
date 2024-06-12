package piscine

func Listmerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	a := n1
	for a.Next != nil {
		a = a.Next
	}
	a.Next = n2

	return n1
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	v := Listmerge(n1, n2)
	a := ListSort(v)
	return a
}
