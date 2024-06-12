package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	x := []int{}
	if l != nil {
		c := l
		for c.Next != nil {
			x = append(x, c.Data)
			c = c.Next
		}
		x = append(x, c.Data)
		i := 1
		for i < len(x) {
			if x[i-1] > x[i] {
				x[i-1], x[i] = x[i], x[i-1]
				i = 1
			} else {
				i++
			}
		}
		i = 1
		l.Data = x[0]
		d := l
		for d.Next != nil {
			d = d.Next
			d.Data = x[i]
			i++
		}
		return l
	} else {
		return nil
	}
}
